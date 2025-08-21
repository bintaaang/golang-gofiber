package routes

import (
	"gofiber-endpoint/handlers"
	middleware "gofiber-endpoint/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes mendefinisikan semua route
func SetupRoutes(app *fiber.App) {
	x := middleware.Load()
	app.Use(middleware.NewIPWhitelistMiddleware(x))
	app.Get("/api/users", handlers.GetAllUsers)
	app.Post("/api/users/post", handlers.CreateDataUser)
	app.Get("/api/users/:id", handlers.GetUserId)
	app.Put("/api/users/:id", handlers.UpdateUser)
	app.Delete("/api/users/:id", handlers.DeleteUserId)
	app.Get("/api/users/getall/ali", handlers.GetAllUsers)
}