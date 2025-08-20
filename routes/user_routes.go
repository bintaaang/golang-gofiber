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
	app.Get("/users", handlers.GetAllUsers)
	app.Post("/users/post", handlers.CreateDataUser)
	app.Get("/users/:id", handlers.GetUserId)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUserId)
	app.Get("/users/getall/ali", handlers.GetAllUsers)

}

