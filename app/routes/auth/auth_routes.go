package auth_routes

import "github.com/gofiber/fiber/v3"

func SetUpAuthRoute(route fiber.Router){
	route.Get("/login", func(c fiber.Ctx) error {
		return c.SendString("/api/auth/login")
	})
	route.Get("/post", func(c fiber.Ctx) error {
		return c.SendString("/api/auth/register")
	})
}