package handlers

import (
	"github.com/gofiber/fiber/v2"
	//"time"
	"gofiber-endpoint/models"
	"gofiber-endpoint/database"
)

func AssignCourier(ctx *fiber.Ctx) error {
    type Request struct {
        PickupID  uint `json:"pickup_id"`
        CourierID uint `json:"courier_id"`
    }

    var body Request
    if err := ctx.BodyParser(&body); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code" : 400,
            "error": err.Error(),
        })
    }

    var pickup models.PickupRequest
    if err := database.UsingPostgre.First(&pickup, body.PickupID).Error; err != nil {
        return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "code": 404,
            "error": "pickup not found",
        })
    }

    pickup.CourierID = &body.CourierID
    pickup.Status = models.PickupStatuses.Assigned
    database.UsingPostgre.Save(&pickup)

    statusLog := models.PackageStatus{
        PickupRequestID: pickup.ID,
        TrackingNo:      pickup.TrackingNo,
        Status:          models.PickupStatuses.Assigned,
        UpdatedByID:     0,
        Note:            "Courier assigned",
    }
    database.UsingPostgre.Create(&statusLog)

    return ctx.JSON(pickup)
}