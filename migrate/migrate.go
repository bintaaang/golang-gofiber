package migrate

import (
	"fmt"
	"gofiber-endpoint/database"
	"gofiber-endpoint/models"
	"log"
)

func MigrateDatabase() {
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("Database migrated successfully!")
}
