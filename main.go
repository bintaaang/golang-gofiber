package main

import (
	"gofiber-endpoint/database"
	//"gofiber-endpoint/migrate"
	"gofiber-endpoint/routes"
	"log"
	//"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitAllDBs()
	//migrate.MigrateDatabase()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
