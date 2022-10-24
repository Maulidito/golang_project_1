package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TotalPrice float64 `form:"total_price"`
	Done       bool
	Cart       []Cart `gorm:"foreignKey:Transaction_Fk" json:"-"`
	User_Fk    int    `form:"user_fk"`
}
