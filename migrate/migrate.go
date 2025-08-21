package migrate

import (
	"fmt"
	"gofiber-endpoint/database"
	"gofiber-endpoint/models"
	"log"
)

func MigrateDatabase() {
	err := database.UsingPostgre.AutoMigrate(
		&models.User{},
        &models.PickupRequest{},
        &models.PackageStatus{},
    )
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("Database migrated successfully!")
}
