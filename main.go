package main

import (
	"github.com/absensi/app/services"
	"github.com/absensi/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	apps := fiber.New(fiber.Config{
		// Views: engine,
	})
	apps.Use(cors.New())
	apps.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	apps.Use(recover.New())
	services := services.New()
	services.Db.GetInstanceConnect()
	defer services.Db.Conn.Close()

	router.Routes(apps, services)

	apps.Listen(":3000")
}
