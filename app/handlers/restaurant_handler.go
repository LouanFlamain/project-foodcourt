package handlers

import (
	"project/foodcourt/store"

	"github.com/gofiber/fiber/v3"
)

func GetAllRestaurant(c fiber.Ctx, store *store.Store) error{
		restaurants, err := store.GetAllRestaurant()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve restaurants",
			})
		}

		err = c.JSON(restaurants)
		return err

}