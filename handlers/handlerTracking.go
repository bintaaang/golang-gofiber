package handlers

import (
	"github.com/gofiber/fiber/v2"
	//"time"
	"gofiber-endpoint/models"
	"gofiber-endpoint/database"
)

func TrackPackage(ctx *fiber.Ctx) error {
    trackingNo := ctx.Params("tracking_no")

    var pickup models.PickupRequest
    if err := database.UsingPostgre.Preload("Courier").First(&pickup, "tracking_no = ?", trackingNo).Error; err != nil {
        return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tracking not found"})
    }

    var logs []models.PackageStatus
    database.UsingPostgre.Where("tracking_no = ?", trackingNo).Order("created_at asc").Find(&logs)

    return ctx.JSON(fiber.Map{
        "pickup": pickup,
        "logs":   logs,
    })
}
