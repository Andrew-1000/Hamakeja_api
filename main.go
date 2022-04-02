package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"hamakeja_api/database"
	"github.com/jinzhu/gorm"
	"hamakeja_api/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gofiber/fiber/v2/middleware/logger"
	
)


func setupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	api.Post("/truck", models.CreateTruck)
	api.Get("/trucks", models.GetTrucks)
	api.Get("/truck/:id", models.GetTruck)
	api.Patch("/truck/:id", models.UpdateTruck)
	api.Delete("/truck/:id", models.DeleteTruck)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("postgres", "host=localhost port=5432 user=andru dbname=hamakeja_api password=1234")

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connection opened to db")

	database.DBConn.AutoMigrate(&models.Truck{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()

	setupRoutes(app)

	app.Listen(":3000")

	defer database.DBConn.Close()
}