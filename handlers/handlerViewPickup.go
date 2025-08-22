package handlers

import (
	"github.com/gofiber/fiber/v2"
	//"time"
	"gofiber-endpoint/models"
	"gofiber-endpoint/database"
)

func GetMyPickups(ctx *fiber.Ctx) error {
    courierID := ctx.Params("courier_id")

    var pickups []models.PickupRequest
    database.UsingPostgre.Where("courier_id = ?", courierID).Find(&pickups)

    return ctx.JSON(pickups)
}