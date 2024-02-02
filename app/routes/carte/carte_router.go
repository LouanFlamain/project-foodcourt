package carte_router


import (
	"foodcourt/app/handlers"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func GetCarteByRestaurantId(route fiber.Router, myStore *stores.Store) {
	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/carte")
	})
	route.Get("/:restaurant_id", func(c fiber.Ctx) error {
		restaurantID, err := c.ParamsInt("restaurant_id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid restaurant_id",
				"success": false,
			})
		}

		return handlers.GetCarteByRestaurantId(c, myStore, restaurantID)

	})
	

}
