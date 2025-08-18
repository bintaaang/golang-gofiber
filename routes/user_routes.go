package routes

import (
	"gofiber-endpoint/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes mendefinisikan semua route
func SetupRoutes(app *fiber.App) {

	app.Get("/users", handlers.GetAllUsers)
	app.Post("/users/post", handlers.CreateDataUser)
	app.Get("/users/:id", handlers.GetUserId)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUserId)
	app.Get("/users/getall/ali", handlers.GetAllUsers)

}

