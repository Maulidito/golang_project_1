//go:build wireinject
// +build wireinject

package container

import (
	"github.com/Maulidito/personal_project_go/controller"
	"github.com/Maulidito/personal_project_go/dataservice"
	"github.com/Maulidito/personal_project_go/middleware"
	"github.com/Maulidito/personal_project_go/models"
	"github.com/Maulidito/personal_project_go/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var userSet = wire.NewSet(
	dataservice.NewDataServiceUser,
	wire.Bind(new(dataservice.IDataService[models.User]), new(*dataservice.DataServiceUser)),
)

var productSet = wire.NewSet(
	dataservice.NewDataServiceProduct,
	wire.Bind(new(dataservice.IDataService[models.Product]), new(*dataservice.DataServiceProduct)),
)

var cartSet = wire.NewSet(
	dataservice.NewDataServiceCart,
	wire.Bind(new(dataservice.IDataService[models.Cart]), new(*dataservice.DataServiceCart)),
)

var transSet = wire.NewSet(
	dataservice.NewDataServiceTransaction,
	wire.Bind(new(dataservice.IDataService[models.Transaction]), new(*dataservice.DataServiceTransaction)),
)

func InitializeAuthMiddleware(db *gorm.DB) *middleware.JwtAuthMiddleware {
	wire.Build(middleware.NewJwtAuthMiddleware, userSet)
	return nil
}

func InitializedServiceAuth(db *gorm.DB) service.IServiceAuth {

	wire.Build(service.NewServiceAuth, userSet)
	return nil
}

func InitializedControllerUser(db *gorm.DB) controller.IController {
	wire.Build(controller.NewControllerUser, service.NewServiceAuth, userSet, InitializeAuthMiddleware)
	return nil
}

func InitializedControllerProduct(db *gorm.DB) controller.IController {
	wire.Build(controller.NewControllerProduct, userSet, productSet, InitializeAuthMiddleware)
	return nil
}

func InitializedControllerCart(db *gorm.DB) controller.IController {
	wire.Build(controller.NewControllerCart, cartSet, InitializeAuthMiddleware)
	return nil
}

func InitializedControllerTrans(db *gorm.DB) controller.IController {
	wire.Build(controller.NewControllerTransaction, cartSet, transSet, InitializeAuthMiddleware)
	return nil
}
