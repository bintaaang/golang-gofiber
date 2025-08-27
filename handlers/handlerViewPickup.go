package handlers

import (
	"github.com/gofiber/fiber/v2"
	//"time"
	"gofiber-endpoint/models"
	"gofiber-endpoint/database"
	//"gofiber-endpoint/validators"
	"gofiber-endpoint/middleware"
)

func GetMyPickups(ctx *fiber.Ctx) error {
    courierID := ctx.Params("courier_id")

    var pickups []models.PickupRequest
    database.UsingPostgre.Where("courier_id = ?", courierID).Find(&pickups)

    return ctx.JSON(pickups)
}
func ViewAllPickup(ctx *fiber.Ctx) error {
	var pickup []models.PickupRequest

	// isTrue := validators.CheckUser("admin"); if isTrue == false{
	// 	return ctx.Status(500).JSON(fiber.Map{
	// 		"error": "invalid role",
	// 	})
	// }
	err := middleware.AuthMiddleware(ctx); if err != nil {
			return err
	}
	role := ctx.Locals("role"); if role != "admin" {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "invlid role",
		})
	}

	err = database.UsingPostgre.Table(`pickup_requests`).Where(`mark = 0`).Find(&pickup).Error; if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"code" : 200,
			"message": "internal error",
		})
	}
	return ctx.Status(200).JSON(pickup)
}