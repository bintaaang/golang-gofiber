package handlers

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"gofiber-endpoint/models"
	"gofiber-endpoint/database"
)

func CreatePickupRequest(ctx *fiber.Ctx) error {
    type Request struct {
        Name        string `json:"name"`
        Phone       string `json:"phone"`
        AddressFrom string `json:"address_from"`
        AddressTo   string `json:"address_to"`
    }

    var body Request
    if err := ctx.BodyParser(&body); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    trackingNo := "TRK" + time.Now().Format("20060102150405")

    request := models.PickupRequest{
        Name:        body.Name,
        Phone:       body.Phone,
        AddressFrom: body.AddressFrom,
        AddressTo:   body.AddressTo,
        TrackingNo:  trackingNo,
        Status:      models.PickupStatuses.Pending,
    }

    database.UsingPostgre.Create(&request)

    return ctx.JSON(fiber.Map{
        "tracking_no": trackingNo,
        "status":      request.Status,
    })
}