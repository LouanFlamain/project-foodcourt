package admin_routes

import "github.com/gofiber/fiber/v3"

func SetUpAdminRoute(group fiber.Router){

	group.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/admin")
	})
}