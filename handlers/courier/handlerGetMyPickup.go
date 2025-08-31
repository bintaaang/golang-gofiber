package courier

import (
	"fmt"
	"gofiber-endpoint/database"
	"gofiber-endpoint/models"

	"github.com/gofiber/fiber/v2"
	//"gofiber-endpoint/middleware"
)

// func GetMyPickups(ctx *fiber.Ctx) error {
// 	courierID := ctx.Params("courier_id")
// 	//claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
// 	// err := middleware.AuthMiddleware(ctx); if err != nil {
// 	// 	return err
// 	// };
// 	coorier_id := ctx.Locals("userID")

// 	var pickups []models.PickupRequest
// 	database.UsingPostgre.Where("courier_id = ?", courierID).Find(&pickups)

// 	return ctx.JSON(pickups)
// }
func GetMyPickups(ctx *fiber.Ctx) error {
    courierIDParam := ctx.Params("courier_id")
    courierIDToken := ctx.Locals("userID")

    if courierIDParam != fmt.Sprint(courierIDToken) {
        return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"code": "200",
            "message": "you are not allowed to access other courier's pickups",
        })
    }

    var pickups []models.PickupRequest
    database.UsingPostgre.Where("courier_id = ?", courierIDParam).Find(&pickups)

    return ctx.JSON(pickups)
}
