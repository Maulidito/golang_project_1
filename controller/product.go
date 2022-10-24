package controller

import (
	"fmt"

	"github.com/Maulidito/personal_project_go/dataservice"
	"github.com/Maulidito/personal_project_go/middleware"
	"github.com/Maulidito/personal_project_go/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ControllerProduct struct {
	productDataService dataservice.IDataService[models.Product]
	userDataService    *dataservice.DataServiceUser
	middlewareAuth     *middleware.JwtAuthMiddleware
	db                 *gorm.DB
}

func (ctrl *ControllerProduct) MountRouter(c fiber.Router) {

	routerGroup := c.Group("/product")

	routerGroup.Get("/", ctrl.GetAll)
	routerGroup.Post("/", ctrl.middlewareAuth.Authentication, ctrl.Add)
	routerGroup.Patch("/:productId", ctrl.middlewareAuth.Authentication, ctrl.Update)
	routerGroup.Delete("/:productId", ctrl.middlewareAuth.Authentication, ctrl.Delete)

}

func NewControllerProduct(dataServiceProd dataservice.IDataService[models.Product], db *gorm.DB, dataServiceUser *dataservice.DataServiceUser, midAuth *middleware.JwtAuthMiddleware) IController {
	return &ControllerProduct{productDataService: dataServiceProd, db: db, userDataService: dataServiceUser, middlewareAuth: midAuth}
}

func (ctrl *ControllerProduct) GetAll(c *fiber.Ctx) error {
	listProduct, err := ctrl.productDataService.GetAll(c, ctrl.db)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Send(nil)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": listProduct,
	})

}

func (ctrl *ControllerProduct) GetOne(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ControllerProduct) Add(c *fiber.Ctx) error {

	//user := c.Locals("user").(*models.User)

	formProduct := models.Product{}
	if err := c.BodyParser(&formProduct); err != nil {
		return fiber.ErrBadRequest
	}

	//fmt.Println(user)
	image, _ := c.FormFile("image")
	fmt.Println(image)

	fmt.Println(formProduct)

	// formProduct.User_Fk = user.ID

	// if err := ctrl.productDataService.Add(c, ctrl.db, &formProduct); err != nil {
	// 	return fiber.ErrInternalServerError
	// }

	return c.SendStatus(fiber.StatusOK)
}

func (ctrl *ControllerProduct) Update(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("productId")

	if err != nil {
		return fiber.ErrBadGateway
	}

	formProduct := models.Product{}
	if err := c.BodyParser(&formProduct); err != nil {
		return fiber.ErrBadRequest
	}

	formProduct.ID = uint(productId)

	if err := ctrl.productDataService.Update(c, ctrl.db, &formProduct); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": formProduct,
	})

}

func (ctrl *ControllerProduct) Delete(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("productId")

	if err != nil {
		return fiber.ErrBadGateway
	}

	err = ctrl.productDataService.Delete(c, ctrl.db, productId)

	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusOK)
}
