package dataservice

import (
	"github.com/Maulidito/personal_project_go/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DataServiceProduct struct {
}

func NewDataServiceProduct() *DataServiceProduct {
	return &DataServiceProduct{}
}

func (dataService *DataServiceProduct) GetAll(ctx *fiber.Ctx, db *gorm.DB) (dataReturn *[]models.Product, err error) {
	db.Find(&dataReturn)
	return
}

func (dataService *DataServiceProduct) GetOne(ctx *fiber.Ctx, db *gorm.DB, id int) (data *models.Product, err error) {
	db.Where("id = ?", id).Find(&data)
	return
}

func (dataService *DataServiceProduct) Add(ctx *fiber.Ctx, db *gorm.DB, data *models.Product) error {
	db.Create(&data)
	return nil
}

func (dataService *DataServiceProduct) Update(ctx *fiber.Ctx, db *gorm.DB, data *models.Product) error {
	db.Where("id = ?", data.ID).Updates(&data)
	return nil
}

func (dataService *DataServiceProduct) Delete(ctx *fiber.Ctx, db *gorm.DB, id int) error {
	db.Where("id = ?", id).Delete(&models.Product{})
	return nil
}
