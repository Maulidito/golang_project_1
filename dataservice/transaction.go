package dataservice

import (
	"github.com/Maulidito/personal_project_go/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DataServiceTransaction struct {
}

func NewDataServiceTransaction() *DataServiceTransaction {
	return &DataServiceTransaction{}
}

func (dataService *DataServiceTransaction) GetAll(ctx *fiber.Ctx, db *gorm.DB) (dataReturn *[]models.Transaction, err error) {
	db.Find(&dataReturn)
	return
}

func (dataService *DataServiceTransaction) GetOne(ctx *fiber.Ctx, db *gorm.DB, id int) (data *models.Transaction, err error) {
	db.Where("user_fk = ?", id).Find(&data)
	return
}

func (dataService *DataServiceTransaction) Add(ctx *fiber.Ctx, db *gorm.DB, data *models.Transaction) error {
	db.Create(&data)
	return nil
}

func (dataService *DataServiceTransaction) Update(ctx *fiber.Ctx, db *gorm.DB, data *models.Transaction) error {
	db.Where("id = ?", data.ID).Updates(&data)
	return nil
}

func (dataService *DataServiceTransaction) Delete(ctx *fiber.Ctx, db *gorm.DB, id int) error {
	db.Where("id = ?", id).Delete(&models.Transaction{})
	return nil
}
