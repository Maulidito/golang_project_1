package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `form:"name"`
	Price       float64 `form:"price"`
	Description string  `form:"description"`
	Quantity    int     `form:"quantity"`
	Image_path  string  `form:"image"`
	User_Fk     uint    `form:"user_fk"`
	Cart        []Cart  `gorm:"foreignKey:Product_Fk" json:"-"`
}
