package main

import (
	"gofiber-endpoint/database"
	//"gofiber-endpoint/migrate"
	"gofiber-endpoint/routes"
	"log"
	//"os"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.Use(cors.New(cors.Config{
        AllowOrigins: "*", // atau spesifik: "http://localhost:3001, https://flask-skripsi.my.id"
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    }))
	
	database.InitAllDBs()
	//migrate.MigrateDatabase()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}