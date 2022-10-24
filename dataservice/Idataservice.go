package dataservice

import (
	"github.com/Maulidito/personal_project_go/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IDataService[T interface {
	models.Cart | models.Product | models.Transaction | models.User
}] interface {
	GetAll(ctx *fiber.Ctx, db *gorm.DB) (*[]T, error)
	GetOne(ctx *fiber.Ctx, db *gorm.DB, id int) (*T, error)
	Add(ctx *fiber.Ctx, db *gorm.DB, data *T) error
	Update(ctx *fiber.Ctx, db *gorm.DB, data *T) error
	Delete(ctx *fiber.Ctx, db *gorm.DB, id int) error
}
