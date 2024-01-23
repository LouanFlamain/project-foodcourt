package seller_routes

import "github.com/gofiber/fiber/v3"

func SetUpSellerRoute(route fiber.Router){
	
	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/seller")
	})
}