package seller_routes

import (
	"encoding/json"
	"foodcourt/app/api/request"
	"foodcourt/app/handlers"
	"foodcourt/app/model"
	"foodcourt/app/stores"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func SetUpSellerRoute(route fiber.Router, myStore *stores.Store) {

	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/seller")
	})

	// Suprimmer la carte des restaurant 
	route.Delete("/carte/:id", func(c fiber.Ctx) error {
		carteID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid restaurant_id",
				"success": false,
			})
		}

		return handlers.DeleteRestaurantById(c, myStore, carteID)
	})

	// creation d'une carte 
	route.Post("/carte/create", func(c fiber.Ctx) error {
		var body request.CreateCarte

		
		err := json.Unmarshal(c.Body(), &body)
		if err != nil {
			// Gérer l'erreur si le JSON est mal formé
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		return handlers.CreateCarte(c, myStore, model.CarteItem{RestaurantId: body.RestaurantId , Description: body.Description , Price: body.Price})

	})

	// creer un produit à la carte 

	route.Post("/product/create", func(c fiber.Ctx) error {
		var body request.CreateProduct

		
		err := json.Unmarshal(c.Body(), &body)
		if err != nil {
			// Gérer l'erreur si le JSON est mal formé
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		return handlers.CreateProduct(c, myStore, model.ProductItem{Produit: body.Produit , Price: body.Price , CarteId: body.CarteId , CategoryId: body.CategoryId })

	})

	// supprimer les produis d'une carte  
	route.Delete("/products/:id", func(c fiber.Ctx) error {
		productID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid restaurant_id",
				"success": false,
			})
		}

		return handlers.DeleteRProductsById(c, myStore, productID)
	})

	// GET commande by id 
	route.Get("/commande/:id", func(c fiber.Ctx) error {
		commandeID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid restaurant_id",
				"success": false,
			})
		}

		return handlers.GetCommandeById(c, myStore, commandeID)
	})
	// GET commande by restaurantId 
	route.Get("/commande/restaurant/:id", func(c fiber.Ctx) error {
		commandeID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid restaurant_id",
				"success": false,
			})
		}

		return handlers.GetAllCommandeByRestaurantId(c, myStore, commandeID)
	})

	// Update commande By Id

	route.Patch("/commande/update/:id", func(c fiber.Ctx) error {
		var body request.UpdateCommande

		err := json.Unmarshal(c.Body(), &body)
		if err != nil {
			// Gérer l'erreur si le JSON est mal formé
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		commandeID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid restaurant_id",
				"success": false,
			})
		}
		return handlers.UpdateCommande(c, myStore, commandeID,model.CommandeItem{State: body.State}, body.Mercure  )

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

}

