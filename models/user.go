package models

import "gorm.io/gorm"

type User struct {
	gorm.Model  `json:"-"`
	Name        string      `form:"name"`
	Password    string      `form:"password"`
	Email       string      `form:"email"`
	Cart        []Cart      `gorm:"foreignKey:User_Fk" json:"-"`
	Product     []Product   `gorm:"foreignKey:User_Fk" json:"-"`
	Transaction Transaction `gorm:"foreignKey:User_Fk" json:"-"`
}
