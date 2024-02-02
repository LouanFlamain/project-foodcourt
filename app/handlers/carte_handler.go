package handlers

import (
	// "encoding/json"
	// "foodcourt/app/api/response"
	// "foodcourt/app/auth"
	// "foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)


func GetCarteByRestaurantId(c fiber.Ctx, carteStore *stores.Store , restaurant_id int) error {

	res, err := carteStore.GetCarteByRestaurantId(restaurant_id)

	if err != nil {
		err = c.JSON(fiber.Map{
			"data": fiber.Map{
				"error":   err,
				"success": false,
			},
		})
		return err
	}
	err = c.JSON(fiber.Map{
		"data": fiber.Map{
			"succes": res,
		},
	})
	return err
}