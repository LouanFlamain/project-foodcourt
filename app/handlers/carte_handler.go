package handlers

import (

	"foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

// trouver le menu d'un restaurant selon l'id
func GetCarteByRestaurantId(c fiber.Ctx, carteStore *stores.Store, restaurant_id int) error {

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

// supprimer le menu d'un restaurant selon l'id
func DeleteRestaurantById(c fiber.Ctx, carteStore *stores.Store, carte_id int) error {

	res, err := carteStore.DeleteCarteById(carte_id)

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

// creer une carte
func CreateCarte(c fiber.Ctx, carteStore *stores.Store, carte model.CarteItem) error {
	_, err := carteStore.CreateCarte(carte)

	if err != nil {
		err = c.JSON(fiber.Map{
			"data": fiber.Map{
				"error":  err,
				"succes": false,
			},
		})
		return err
	}
	err = c.JSON(fiber.Map{
		"data": fiber.Map{
			"succes": true,
		},
	})
	return err
}

// get les produits d'une carte
func GetProductsByCarteId(c fiber.Ctx, ProductStore *stores.Store, carte_id int) error {

	res, err := ProductStore.GetProductsByCarteId(carte_id)

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

// Ajouter une produit à la carte

func CreateProduct(c fiber.Ctx, ProductStore *stores.Store, product model.ProductItem) error {
	_, err := ProductStore.CreateProduct(product)

	if err != nil {
		err = c.JSON(fiber.Map{
			"data": fiber.Map{
				"error":  err,
				"succes": false,
			},
		})
		return err
	}
	err = c.JSON(fiber.Map{
		"data": fiber.Map{
			"succes": true,
		},
	})
	return err

}

// supprimer un produits
func DeleteRProductsById(c fiber.Ctx, ProductStore *stores.Store, carte_id int) error {

	res, err := ProductStore.DeleteProductById(carte_id)

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
