package customer_routes

import "github.com/gofiber/fiber/v3"

func SetUpCustomerRoute(group fiber.Router){
	group.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/customer")
	})
}