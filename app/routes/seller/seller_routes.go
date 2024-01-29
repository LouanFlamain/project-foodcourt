package seller_routes

import (
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func SetUpSellerRoute(route fiber.Router, myStore *stores.Store) {

	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/seller")
	})
}
