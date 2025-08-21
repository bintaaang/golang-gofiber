package handlers

import (
	"github.com/gofiber/fiber/v2"
	//"time"
	"gofiber-endpoint/models"
	"gofiber-endpoint/database"
)

func TrackPackage(c *fiber.Ctx) error {
    trackingNo := c.Params("tracking_no")

    var pickup models.PickupRequest
    if err := database.UsingPostgre.Preload("Courier").First(&pickup, "tracking_no = ?", trackingNo).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tracking not found"})
    }

    var logs []models.PackageStatus
    database.UsingPostgre.Where("tracking_no = ?", trackingNo).Order("created_at asc").Find(&logs)

    return c.JSON(fiber.Map{
        "pickup": pickup,
        "logs":   logs,
    })
}
