package main

import (
	"gofiber-endpoint/database"
	//"gofiber-endpoint/migrate"
	"gofiber-endpoint/routes"
	"log"
	"os"
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
	port := os.Getenv("PORTX")
	if port == "" {
		log.Fatal("port is not in environment")
	}
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Fatal("domain is not in environment")
	}
	app.Use(cors.New(cors.Config{
        AllowOrigins: "*", // atau spesifik: "http://localhost:3001, https://flask-skripsi.my.id"
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    }))
	
	database.InitAllDBs()
	//migrate.MigrateDatabase()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(domain + port))
	//app.Listen(":3000")
}