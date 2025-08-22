package handlers

import (
	"github.com/gofiber/fiber/v2"
	//"time"
	"gofiber-endpoint/models"
	"gofiber-endpoint/database"
)

func UpdatePickupStatus(ctx *fiber.Ctx) error {
    type Request struct {
        PickupID  uint   `json:"pickup_id"`
        Status    string `json:"status"`
        Note      string `json:"note"`
        UpdatedBy uint   `json:"updated_by"` // user/admin/kurir ID
    }

    var body Request
    if err := ctx.BodyParser(&body); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    var pickup models.PickupRequest
    if err := database.UsingPostgre.First(&pickup, body.PickupID).Error; err != nil {
        return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "pickup not found"})
    }

    pickup.Status = body.Status
    database.UsingPostgre.Save(&pickup)

    statusLog := models.PackageStatus{
        PickupRequestID: pickup.ID,
        TrackingNo:      pickup.TrackingNo,
        Status:          body.Status,
        UpdatedByID:     body.UpdatedBy,
        Note:            body.Note,
    }
    database.UsingPostgre.Create(&statusLog)

    return ctx.JSON(pickup)
}