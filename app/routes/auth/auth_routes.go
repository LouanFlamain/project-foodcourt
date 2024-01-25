package auth_routes

import (
	"project/foodcourt/store"

	"github.com/gofiber/fiber/v3"
)

func SetUpAuthRoute(route fiber.Router, myStore *store.Store){
	route.Get("/login", func(c fiber.Ctx) error {
		return c.SendString("/api/auth/login")
	})
	route.Get("/post", func(c fiber.Ctx) error {
		return c.SendString("/api/auth/register")
	})
}