package handlers

import (
	"github.com/gofiber/fiber/v2"
	//"time"
	"gofiber-endpoint/models"
	"gofiber-endpoint/database"
)

func GetMyPickups(c *fiber.Ctx) error {
    courierID := c.Params("courier_id") // contoh dari JWT atau param

    var pickups []models.PickupRequest
    database.UsingPostgre.Where("courier_id = ?", courierID).Find(&pickups)

    return c.JSON(pickups)
}