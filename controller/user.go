package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/Maulidito/personal_project_go/dataservice"
	"github.com/Maulidito/personal_project_go/middleware"
	"github.com/Maulidito/personal_project_go/models"
	"github.com/Maulidito/personal_project_go/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type ControllerUser struct {
	userDataService *dataservice.DataServiceUser
	db              *gorm.DB
	userAuthService service.IServiceAuth
	middlewareAuth  *middleware.JwtAuthMiddleware
}

func (ctrl *ControllerUser) MountRouter(c fiber.Router) {
	routerGroup := c.Group("/user")
	routerGroup.Get("/", ctrl.GetAll)
	routerGroup.Post("/login", ctrl.Login)
	routerGroup.Post("/register", ctrl.Add)
	routerGroup.Patch("/", ctrl.middlewareAuth.Authentication, ctrl.Update)
	routerGroup.Delete("/", ctrl.middlewareAuth.Authentication, ctrl.Delete)

}

func NewControllerUser(dataServiceProd *dataservice.DataServiceUser, db *gorm.DB, authService service.IServiceAuth, midAuth *middleware.JwtAuthMiddleware) IController {
	return &ControllerUser{userDataService: dataServiceProd, db: db, userAuthService: authService, middlewareAuth: midAuth}
}

func (ctrl *ControllerUser) GetAll(c *fiber.Ctx) error {
	listProduct, err := ctrl.userDataService.GetAll(c, ctrl.db)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Send(nil)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": listProduct,
	})

}

func (ctrl *ControllerUser) GetOne(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ControllerUser) Delete(c *fiber.Ctx) error {

	user := c.Locals("user").(*models.User)

	err := ctrl.userDataService.Delete(c, ctrl.db, int(user.ID))

	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})

}

func (ctrl *ControllerUser) Add(c *fiber.Ctx) error {
	formUser := models.User{}
	if err := c.BodyParser(&formUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}
	if err := ctrl.userAuthService.Register(c, &formUser); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (ctrl *ControllerUser) Update(c *fiber.Ctx) error {
	formUser := models.User{}
	if err := c.BodyParser(&formUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}
	user := c.Locals("name").(*models.User)

	formUser.ID = user.ID

	if err := ctrl.userDataService.Update(c, ctrl.db, &formUser); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})

}

func (ctrl *ControllerUser) Login(c *fiber.Ctx) error {
	formUser := models.User{}
	if err := c.BodyParser(&formUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}

	if err := ctrl.userAuthService.Login(c, formUser.Name, formUser.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		middleware.JwtCustomClaim{
			Name: formUser.Name,
			RegisteredClaims: &jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 500)),
			},
		},
	)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_AUTH")))

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{"Token": tokenString})

	return nil
}
