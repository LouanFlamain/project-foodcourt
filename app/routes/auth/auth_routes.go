package auth_routes

import (
	"foodcourt/app/handlers"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func SetUpAuthRoute(route fiber.Router, myStore *stores.Store) {
	route.Post("/register", func(c *fiber.Ctx) error {
		return handlers.RegisterHandler(c, myStore.UserInterface)
	})

	route.Post("/login", func(c *fiber.Ctx) error {
		return handlers.LoginHandler(c, myStore.UserInterface)
	})

}
