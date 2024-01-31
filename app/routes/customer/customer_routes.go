package customer_routes

import (
	"encoding/json"
	"foodcourt/app/api/request"
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
	route.Post("/restaurant/create", func(c fiber.Ctx)error{
		var body request.CreateRestaurantRequestType
		err := json.Unmarshal(c.Body(), &body)
		if err != nil {
			// Gérer l'erreur si le JSON est mal formé
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		return handlers.CreateRestaurant(c, myStore, body)
	})

	//restaurant

}
