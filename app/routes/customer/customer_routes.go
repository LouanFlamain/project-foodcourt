package customer_routes

import (
	"project/foodcourt/store"

	"github.com/gofiber/fiber/v3"
)

func SetUpCustomerRoute(route fiber.Router, myStore *store.Store){
	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/customer")
	})

	//restaurant
	
}