package admin_routes

import (
	"project/foodcourt/handlers"
	"project/foodcourt/store"

	"github.com/gofiber/fiber/v3"
)

func SetUpAdminRoute(route fiber.Router, myStore *store.Store){

	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/admin")
	})
	route.Get("/restaurant/all", func(c fiber.Ctx) error {
		return handlers.GetAllRestaurant(c, myStore)
	})
}