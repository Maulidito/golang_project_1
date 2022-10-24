package main

import (
	"log"

	"github.com/Maulidito/personal_project_go/app/database"
	"github.com/Maulidito/personal_project_go/controller"
	"github.com/Maulidito/personal_project_go/dataservice"
	"github.com/Maulidito/personal_project_go/middleware"
	"github.com/Maulidito/personal_project_go/models"
	"github.com/Maulidito/personal_project_go/service"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db, err := database.NewDatabasePostgres()

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(models.User{}, models.Product{}, models.Transaction{}, models.Cart{})

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	cartDataService := dataservice.NewDataServiceCart()
	productDataService := dataservice.NewDataServiceProduct()
	transactionDataService := dataservice.NewDataServiceTransaction()
	userDataService := dataservice.NewDataServiceUser()

	middlewareAuth := middleware.NewJwtAuthMiddleware(userDataService, db)

	serviceAuth := service.NewServiceAuth(userDataService, db)

	controllerProduct := controller.NewControllerProduct(productDataService, db, userDataService, middlewareAuth)
	controllerUser := controller.NewControllerUser(userDataService, db, serviceAuth, middlewareAuth)
	controllerTransaction := controller.NewControllerTransaction(transactionDataService, db, middlewareAuth, cartDataService)
	controllerCart := controller.NewControllerCart(cartDataService, db, middlewareAuth)

	appGroup := app.Group("/api")

	controllerUser.MountRouter(appGroup)
	controllerProduct.MountRouter(appGroup)
	controllerTransaction.MountRouter(appGroup)
	controllerCart.MountRouter(appGroup)

	app.Listen(":8080")

}
