package handlers

import (
	"github.com/gofiber/fiber/v2"
	//"time"
	"gofiber-endpoint/models"
	"gofiber-endpoint/database"
)

func AssignCourier(c *fiber.Ctx) error {
    type Request struct {
        PickupID  uint `json:"pickup_id"`
        CourierID uint `json:"courier_id"`
    }

    var body Request
    if err := c.BodyParser(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    var pickup models.PickupRequest
    if err := database.UsingPostgre.First(&pickup, body.PickupID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "pickup not found"})
    }

    pickup.CourierID = &body.CourierID
    pickup.Status = models.PickupStatuses.Assigned
    database.UsingPostgre.Save(&pickup)

    // log status
    statusLog := models.PackageStatus{
        PickupRequestID: pickup.ID,
        TrackingNo:      pickup.TrackingNo,
        Status:          models.PickupStatuses.Assigned,
        UpdatedByID:     0, // admin ID, bisa dari JWT
        Note:            "Courier assigned",
    }
    database.UsingPostgre.Create(&statusLog)

    return c.JSON(pickup)
}