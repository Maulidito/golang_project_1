package dataservice

import (
	"fmt"

	"github.com/Maulidito/personal_project_go/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DataServiceCart struct {
}

func NewDataServiceCart() *DataServiceCart {
	return &DataServiceCart{}
}

func (dataService *DataServiceCart) GetAll(ctx *fiber.Ctx, db *gorm.DB) (dataReturn *[]models.Cart, err error) {
	db.Find(&dataReturn)
	return

}

func (dataService *DataServiceCart) GetOne(ctx *fiber.Ctx, db *gorm.DB, id int) (data *models.Cart, err error) {
	db.Where("id = ?", id).Find(&data)
	return
}

func (dataService *DataServiceCart) Add(ctx *fiber.Ctx, db *gorm.DB, data *models.Cart) error {
	db.Create(&data)
	return nil
}

func (dataService *DataServiceCart) Update(ctx *fiber.Ctx, db *gorm.DB, data *models.Cart) error {
	db.Where("id = ?", data.ID).Updates(&data)
	return nil
}

func (dataService *DataServiceCart) Delete(ctx *fiber.Ctx, db *gorm.DB, id int) error {
	db.Where("id = ?", id).Delete(&models.Cart{})
	return nil
}

func (dataService *DataServiceCart) GetTotalCartByUser(ctx *fiber.Ctx, db *gorm.DB, id int) (dataReturn int, err error) {

	rows, err := db.Raw("SELECT DISTINCT sum(p.price * c.quantity) FROM carts AS c INNER JOIN products AS p ON c.product_fk = p.id WHERE c.user_fk = ? GROUP BY c.user_fk", id).Rows()
	fmt.Println(err)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&dataReturn)

	}
	return
}

func (dataService *DataServiceCart) UpdateAll(ctx *fiber.Ctx, db *gorm.DB, id_user int, cart *models.Cart) error {
	_, err := db.Where("user_fk = ?", id_user).Updates(&cart).Rows()

	if err != nil {
		return err
	}
	return nil
}
