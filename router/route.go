package router

import (
	"github.com/absensi/app/handlers"
	"github.com/absensi/app/services"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, services *services.Services) {

	route := app.Group("/api/")

	handlers := &handlers.Handler{
		Ser: services,
	}

	route.Post("/login", handlers.Login)
	// auten := &authorize.AuthorizeUser{
	// 	Ser: services,
	// }
	// routeAuth := app.Group("/api/v1/", auten.Auth())
	// routeAuth.Get("/getuser", handlers.UserLog)
}
