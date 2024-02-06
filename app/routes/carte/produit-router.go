package carte_router


import (
	"foodcourt/app/handlers"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func GetPoduitByRestaurantId(route fiber.Router, myStore *stores.Store) {
	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/produit")
	})
	route.Get("/:carte_id", func(c fiber.Ctx) error {
		carteId, err := c.ParamsInt("carte_id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid carte_id",
				"success": false,
			})
		}

		return handlers.GetCarteByRestaurantId(c, myStore, carteId)

	})	

}

func DeleteProduitById(route fiber.Router, myStore *stores.Store) {
	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("/api/produit")
	})
	route.Get("/:produit_id", func(c fiber.Ctx) error {
		produitId, err := c.ParamsInt("produit_id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid produit_id",
				"success": false,
			})
		}

		return handlers.GetCarteByRestaurantId(c, myStore, produitId)

	})
}
