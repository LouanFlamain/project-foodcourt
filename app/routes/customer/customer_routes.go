package customer_routes

import (
	"foodcourt/app/handlers"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func SetUpCustomerRoute(route fiber.Router, myStore *stores.Store) {
	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/customer")
	})
	route.Get("/restaurant/all", func(c fiber.Ctx) error {
		return handlers.GetAllOpenRestaurant(c, myStore)
	})

	//restaurant

}
