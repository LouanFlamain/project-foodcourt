package customer_routes

import (
	"encoding/json"
	"foodcourt/app/api/request"
	"foodcourt/app/handlers"
	"foodcourt/app/model"
	"foodcourt/app/stores"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func SetUpCustomerRoute(route fiber.Router, myStore *stores.Store) {
	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/customer")
	})

	route.Get("/restaurant/all", func(c fiber.Ctx) error {
		return handlers.GetAllOpenRestaurant(c, myStore)
	})


	// On récupére la carte des restaurant 
	route.Get("/carte/:id", func(c fiber.Ctx) error {
		restaurantID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid restaurant_id",
				"success": false,
			})
		}

		return handlers.GetCarteByRestaurantId(c, myStore, restaurantID)
	})

	// trouver les produis d'une carte  
	route.Get("/products/:id", func(c fiber.Ctx) error {
		carteID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid restaurant_id",
				"success": false,
			})
		}

		return handlers.GetProductsByCarteId(c, myStore, carteID)
	})

	// creer une commande 
	route.Post("/commande/create", func(c fiber.Ctx) error {
		var body request.CreateCommande

		err := json.Unmarshal(c.Body(), &body)
		if err != nil {
			// Gérer l'erreur si le JSON est mal formé
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
	
		content := make([]interface{}, len(body.Content))
		copy(content, body.Content)
	

		return handlers.CreateCommande(c, myStore, model.CommandeItem{Date: body.Date , UserId: body.UserId ,
			 RestaurantId: body.RestaurantId , Content: content , Commentaire: body.Commentaire , State: body.State})

	})

}
