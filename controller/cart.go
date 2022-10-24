package controller

import (
	"github.com/Maulidito/personal_project_go/dataservice"
	"github.com/Maulidito/personal_project_go/middleware"
	"github.com/Maulidito/personal_project_go/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ControllerCart struct {
	cartDataService dataservice.IDataService[models.Cart]
	db              *gorm.DB
	middlewareAuth  *middleware.JwtAuthMiddleware
}

func (ctrl *ControllerCart) MountRouter(c fiber.Router) {
	routerGroup := c.Group("/cart")
	routerGroup.Get("/", ctrl.GetAll)
	routerGroup.Post("/", ctrl.middlewareAuth.Authentication, ctrl.Add)
	routerGroup.Patch("/:cartId", ctrl.middlewareAuth.Authentication, ctrl.Update)
	routerGroup.Delete("/:cartId", ctrl.middlewareAuth.Authentication, ctrl.Delete)

}

func NewControllerCart(dataServiceProd dataservice.IDataService[models.Cart], db *gorm.DB, midAuth *middleware.JwtAuthMiddleware) IController {
	return &ControllerCart{cartDataService: dataServiceProd, db: db, middlewareAuth: midAuth}
}

func (ctrl *ControllerCart) GetAll(c *fiber.Ctx) error {
	listProduct, err := ctrl.cartDataService.GetAll(c, ctrl.db)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Send(nil)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": listProduct,
	})

}

func (ctrl *ControllerCart) GetOne(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ControllerCart) Add(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	formCart := models.Cart{}
	if err := c.BodyParser(&formCart); err != nil {
		return c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}

	formCart.User_Fk = user.ID

	if err := ctrl.cartDataService.Add(c, ctrl.db, &formCart); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusOK)

}

func (ctrl *ControllerCart) Update(c *fiber.Ctx) error {
	cartId, err := c.ParamsInt("cartId")

	if err != nil {
		return fiber.ErrBadGateway
	}
	user := c.Locals("user").(*models.User)
	formCart := models.Cart{}
	if err := c.BodyParser(&formCart); err != nil {
		return c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}

	formCart.ID = uint(cartId)
	formCart.User_Fk = user.ID

	if err := ctrl.cartDataService.Update(c, ctrl.db, &formCart); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusOK)

}

func (ctrl *ControllerCart) Delete(c *fiber.Ctx) error {
	cartId, err := c.ParamsInt("cartId")

	if err != nil {
		return fiber.ErrBadGateway
	}

	if err := ctrl.cartDataService.Delete(c, ctrl.db, cartId); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusOK)

}
