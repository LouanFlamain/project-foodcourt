package admin_routes

import (
	"encoding/json"
	"fmt"
	"project/foodcourt/handlers"
	"project/foodcourt/store"
	"project/foodcourt/structure"
	"project/foodcourt/web/api/request"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func SetUpAdminRoute(route fiber.Router, myStore *store.Store){

	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/admin")
	})


	//------restaurant------


	//récupère TOUT les restaurants
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
		return handlers.CreateNewRestaurantCategory(c, myStore, structure.RestaurantCategoryItem{Name: body.Name})

	})
}