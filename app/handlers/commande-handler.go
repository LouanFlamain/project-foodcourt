package handlers

import (
	"encoding/json"
	"fmt"
	"foodcourt/app/api/mercure"
	"foodcourt/app/model"
	"foodcourt/app/stores"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// creer une commande

func CreateCommande(c fiber.Ctx , commandeStore *stores.Store , commande model.CommandeItem) error {
	_, err := commandeStore.CreateCommande(commande)
	fmt.Println("ceci est l'erreur", err)

	if err != nil {
      err = c.JSON(fiber.Map{
		"data" : fiber.Map{
			"error" : err,
			"succes" : false,
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

// GET commande by id
func GetCommandeById(c fiber.Ctx, CommandeStore *stores.Store, id int) error {

	res, err := CommandeStore.GetCommandeById(id)

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
		"data": res,
	})
	
	return err
}

// Get commandes by restaurantId

func GetAllCommandeByRestaurantId(c fiber.Ctx, CommandeStore *stores.Store, id int) error {

	res, err := CommandeStore.GetAllCommandeByRestaurantId(id)

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
		"data": res,
	})
	return err
}
// Get commandes by restaurantId

func GetAllCommandeByUserId(c fiber.Ctx, CommandeStore *stores.Store, id int) error {

	res, err := CommandeStore.GetAllCommandeByUserId(id)

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
		"data": res,
	})
	return err
}

// Update Commande 
func UpdateCommande(c fiber.Ctx, CommandeStore *stores.Store, id int , commande model.CommandeItem, mercureToken string) error {
     
	_, err := CommandeStore.UpdateCommande(id , commande)

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
			"succes": true,
		},
	})
	res, err := CommandeStore.GetCommandeById(id);
	if err != nil {
		err = c.JSON(fiber.Map{
			"data": fiber.Map{
				"error":   err,
				"success": false,
			},
		})
		return err
	};
	resJSON, err := json.Marshal(res)
	if err != nil {
		log.Println("Erreur lors de la sérialisation de la réponse :", err)
		return err
	}
	log.Println(string(resJSON), mercureToken)

	err = mercure.PublishUpdate(strconv.Itoa(res.UserId), string(resJSON), mercureToken)
	return err
}

