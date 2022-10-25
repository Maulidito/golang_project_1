package main

import (
	"log"

	"github.com/Maulidito/personal_project_go/app/database"
	"github.com/Maulidito/personal_project_go/container"
	"github.com/Maulidito/personal_project_go/models"
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

	controllerProduct := container.InitializedControllerProduct(db)
	controllerUser := container.InitializedControllerUser(db)
	controllerTransaction := container.InitializedControllerTrans(db)
	controllerCart := container.InitializedControllerCart(db)

	appGroup := app.Group("/api")

	controllerUser.MountRouter(appGroup)
	controllerProduct.MountRouter(appGroup)
	controllerTransaction.MountRouter(appGroup)
	controllerCart.MountRouter(appGroup)

	app.Listen(":8080")

}
