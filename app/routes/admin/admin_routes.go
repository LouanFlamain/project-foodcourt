package admin_routes

import "github.com/gofiber/fiber/v3"

func SetUpAdminRoute(route fiber.Router){

	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/admin")
	})
}