package controller

import (
	"github.com/Maulidito/personal_project_go/dataservice"
	"github.com/Maulidito/personal_project_go/middleware"
	"github.com/Maulidito/personal_project_go/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ControllerTransaction struct {
	transDataService dataservice.IDataService[models.Transaction]
	cartDataService  *dataservice.DataServiceCart
	db               *gorm.DB
	middlewareAuth   *middleware.JwtAuthMiddleware
}

func (ctrl *ControllerTransaction) MountRouter(c fiber.Router) {
	routerGroup := c.Group("/transaction")
	routerGroup.Get("/", ctrl.GetAll)
	routerGroup.Post("/", ctrl.middlewareAuth.Authentication, ctrl.Add)
	routerGroup.Patch("/:transId", ctrl.middlewareAuth.Authentication, ctrl.Update)

}

func NewControllerTransaction(dataServiceProd dataservice.IDataService[models.Transaction], db *gorm.DB, midAuth *middleware.JwtAuthMiddleware, cartDataService *dataservice.DataServiceCart) IController {
	return &ControllerTransaction{transDataService: dataServiceProd, db: db, middlewareAuth: midAuth, cartDataService: cartDataService}

}

func (ctrl *ControllerTransaction) GetAll(c *fiber.Ctx) error {
	listProduct, err := ctrl.transDataService.GetAll(c, ctrl.db)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Send(nil)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": listProduct,
	})

}

func (ctrl *ControllerTransaction) GetOne(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ControllerTransaction) Add(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	TotalPrice, _ := ctrl.cartDataService.GetTotalCartByUser(c, ctrl.db, int(user.ID))
	transaction := models.Transaction{TotalPrice: float64(TotalPrice), User_Fk: int(user.ID)}

	if err := ctrl.transDataService.Add(c, ctrl.db, &transaction); err != nil {
		return fiber.ErrInternalServerError
	}

	cart := models.Cart{Transaction_Fk: &transaction.ID}
	if err := ctrl.cartDataService.UpdateAll(c, ctrl.db, int(user.ID), &cart); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusOK)
}

func (ctrl *ControllerTransaction) Update(c *fiber.Ctx) error {
	transId, err := c.ParamsInt("transId")

	if err != nil {
		return fiber.ErrBadGateway
	}

	user := c.Locals("user").(*models.User)

	trans := models.Transaction{User_Fk: int(user.ID), Done: true}

	trans.ID = uint(transId)

	if err := ctrl.transDataService.Update(c, ctrl.db, &trans); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusOK)
}
