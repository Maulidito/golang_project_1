package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Quantity       int   `form:"quantity"`
	User_Fk        uint  `form:"user_fk"`
	Transaction_Fk *uint `form:"transaction_fk" gorm:"type:null"`
	Product_Fk     uint  `form:"product_fk"`
}
