package admin_routes

import (
	"encoding/json"
	"fmt"
	"foodcourt/app/api/request"
	"foodcourt/app/handlers"
	"foodcourt/app/model"
	"foodcourt/app/stores"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func SetUpAdminRoute(route fiber.Router, myStore *stores.Store) {

	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/admin")
	})

	//------restaurant------

	//récupère TOUT les restaurants ouvert ou fermé
	route.Get("/restaurant/all", func(c fiber.Ctx) error {
		return handlers.GetAllRestaurant(c, myStore)
	})

	//supprimer un restaurant
	route.Delete("/restaurant/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			fmt.Println("erreur durant la conversion string -> id")
		}
		return handlers.DeleteRestaurant(c, myStore, id)
	})

	//récupère tout les restaurant en attente de validation d'un admin pour exister
	route.Get("/restaurant/draft", func(c fiber.Ctx) error {
		return handlers.GetDraftRestaurant(c, myStore)
	})

	//accepter le restaurant au sein du foodcourt
	route.Patch("/restaurant/draft/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			fmt.Println("erreur durant la conversion string -> id")
		}
		return handlers.PatchDraftRestaurant(c, myStore, id)
	})

	//créer une nouvelle catégorie de restaurant
	route.Post("/restaurant/category", func(c fiber.Ctx) error {

		var body request.CreateCategoryRequestType

		//parse le body

		err := json.Unmarshal(c.Body(), &body)
		if err != nil {
			// Gérer l'erreur si le JSON est mal formé
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		return handlers.CreateNewRestaurantCategory(c, myStore, model.RestaurantCategoryItem{Name: body.Name})

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
		return handlers.UpdateCommande(c, myStore, commandeID,model.CommandeItem{State: body.State}  )

	})


}
