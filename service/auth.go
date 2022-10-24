package service

import (
	"github.com/Maulidito/personal_project_go/dataservice"
	"github.com/Maulidito/personal_project_go/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IServiceAuth interface {
	Login(ctx *fiber.Ctx, username string, password string) error
	Register(ctx *fiber.Ctx, data *models.User) error
	Logout(ctx *fiber.Ctx) error
}

type ServiceAuth struct {
	dataServiceUser *dataservice.DataServiceUser

	db *gorm.DB
}

func NewServiceAuth(dataService dataservice.IDataService[models.User], db *gorm.DB) IServiceAuth {
	return &ServiceAuth{dataServiceUser: dataService.(*dataservice.DataServiceUser), db: db}
}

func (serv *ServiceAuth) Login(ctx *fiber.Ctx, username string, password string) error {

	dataUser, err := serv.dataServiceUser.GetOneByName(ctx, serv.db, username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func (serv *ServiceAuth) Register(ctx *fiber.Ctx, data *models.User) error {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	data.Password = string(hashPass)

	err = serv.dataServiceUser.Add(ctx, serv.db, data)

	if err != nil {
		return err
	}

	return nil
}

func (serv *ServiceAuth) Logout(ctx *fiber.Ctx) error {

	panic("not implemented") // TODO: Implement
}
