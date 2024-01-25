package handlers

import (
	"fmt"
	"project/foodcourt/store"
	"project/foodcourt/structure"

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

func DeleteRestaurant(c fiber.Ctx, store *store.Store, id int) error{
	err := store.DeleteRestaurant(id)

	if err != nil {
		err := c.JSON(fiber.Map{
			"success" : false,
			"erreur" : err,
		})
		return err
	}

	err = c.JSON(fiber.Map{
		"success" : true,
	})

	return err

}

//new category restaurant

func CreateNewRestaurantCategory(c fiber.Ctx, store *store.Store, item structure.RestaurantCategoryItem) error{
	//var item structure.RestaurantCategoryItem
	err := store.CreateCategory(item)

	if err != nil {
		 err = c.JSON(fiber.Map{
			"success" : false,
			"error" : err,
		})
		return err
	}
	fmt.Println(err)
	return nil
}