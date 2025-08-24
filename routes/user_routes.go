package routes

import (
	"gofiber-endpoint/handlers"
	middleware "gofiber-endpoint/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	x := middleware.Load()
	app.Use(middleware.NewIPWhitelistMiddleware(x))
	//app.Get("/api/users", handlers.GetAllUsers)
	//app.Post("/api/users/post", handlers.CreateDataUser)
	//app.Get("/api/usersid/:id", handlers.GetUserId)
	//app.Put("/api/usersupd/:id", handlers.UpdateUser)
	//app.Delete("/api/usersdel/:id", c.DeleteUserId)
	//app.Get("/api/users/getall/ali", handlers.GetAllUsers)
	app.Post("/api/pickup", handlers.CreatePickupRequest)
    app.Get("/api/track/:tracking_no", handlers.TrackPackage)
    app.Post("/api/assign-courier", handlers.AssignCourier)
    app.Post("/api/update-status", handlers.UpdatePickupStatus)
    app.Get("/api/my-pickups/:courier_id", handlers.GetMyPickups)
	app.Post("/api/register", handlers.Register)
	app.Post("/api/login", handlers.Login)
	protected := app.Group("/", middleware.Protected())
	protected.Get("/api/getallpickup", handlers.ViewAllPickup)

}