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

	//créer une nouvelle catégorie de produits
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
}
