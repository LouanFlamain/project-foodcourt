package seller_routes

import "github.com/gofiber/fiber/v3"

func SetUpSellerRoute(group fiber.Router){
	
	group.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/seller")
	})
}