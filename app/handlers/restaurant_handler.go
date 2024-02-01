package handlers

import (
	"fmt"
	"foodcourt/app/api/request"
	"foodcourt/app/api/response"
	"foodcourt/app/auth"
	"foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

//----------------------------------------------------customer----------------------------------------------------------

func GetAllOpenRestaurant(c fiber.Ctx, stores *stores.Store) error {
	restaurants, err := stores.GetAllOpenRestaurant()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": fiber.Map{
				"success": false,
				"error":   "Failed to retrieve restaurants",
			},
		})
	}

	responseRestaurant := make([]response.GetAllRestaurantResponseType, len(restaurants))
	for i, restaurant := range restaurants {
		responseRestaurant[i] = response.GetAllRestaurantResponseType{
			Id:          restaurant.Id,
			Name:        restaurant.Name,
			Email:       restaurant.Email,
			Picture:     restaurant.Picture,
			Description: restaurant.Description,
			CategoryId:  restaurant.CategoryId,
			Open:        restaurant.Open,
			UserId: restaurant.UserId,
		}
	}

	err = c.JSON(fiber.Map{
		"data": responseRestaurant,
	})
	return err

}

func CreateRestaurant(c fiber.Ctx, store *stores.Store, body request.CreateRestaurantRequestType)error {


	//vérifie que l'utilisateur n'a pas déjà un compte enregistré avec le même mail
	err := store.VerifyUserByMail(body.Email)
	fmt.Println(err)
	if err == nil {
		err := c.JSON(fiber.Map{
			"data" : fiber.Map{
				"success" : false,
				"error" : "email is already use",
			},
		})
		return err
	}
	//création d'un utilisateur de type restaurateur
	hashedPassword, err := auth.HashPassword(body.Password)
	newUser := model.UserItem{
		Username: body.Username,
		Password: hashedPassword,
		Email: body.Email,
		Roles: 1,
	}
	id, err := store.AddRestaurateur(newUser)
	if err != nil {
		err := c.JSON(fiber.Map{
			"data" : fiber.Map{
				"success" : false,
				"error" : err,
			},
		})
		return err
	}
	newRestaurant := model.RestaurantItem{
		Name: body.Name,
		Email: body.Email,
		Description: body.Description,
		CategoryId: body.CategoryId,
		UserId: id,
	}
	err = store.CreateRestaurant(newRestaurant)

	if err != nil{
		err = c.JSON(fiber.Map{
			"data" : fiber.Map{
				"success" : false,
				"error" : err,
			},
		})
	}

	return c.JSON(fiber.Map{
		"data" : fiber.Map{
			"message" : "Restaurant crée avec succès !",
		},
	})
}

//----------------------------------------------------seller------------------------------------------------------------

//----------------------------------------------------admin-------------------------------------------------------------

func GetAllRestaurant(c fiber.Ctx, stores *stores.Store) error {
	restaurants, err := stores.GetAllRestaurant()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": fiber.Map{
				"success": false,
				"error":   "Failed to retrieve restaurants",
			},
		})
	}

	responseRestaurant := make([]response.GetAllRestaurantResponseType, len(restaurants))
	for i, restaurant := range restaurants {
		responseRestaurant[i] = response.GetAllRestaurantResponseType{
			Id:          restaurant.Id,
			Name:        restaurant.Name,
			Email:       restaurant.Email,
			Picture:     restaurant.Picture,
			Description: restaurant.Description,
			CategoryId:  restaurant.CategoryId,
			Open:        restaurant.Open,
		}
	}

	err = c.JSON(fiber.Map{
		"data": responseRestaurant,
	})
	return err

}
func GetDraftRestaurant(c fiber.Ctx, stores *stores.Store) error {
	restaurants, err := stores.GetDraftRestaurant()

	if err != nil {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"success": false,
				"error":   err,
			},
		})
	}
	responseRestaurant := make([]response.GetDraftRestaurantResponseType, len(restaurants))
	for i, restaurant := range restaurants {
		responseRestaurant[i] = response.GetDraftRestaurantResponseType{
			Id:          restaurant.Id,
			Name:        restaurant.Name,
			Email:       restaurant.Email,
			Picture:     restaurant.Picture,
			Description: restaurant.Description,
			CategoryId:  restaurant.CategoryId,
			Draft:       restaurant.Draft,
		}
	}

	err = c.JSON(fiber.Map{
		"data": responseRestaurant,
	})
	return err

}

func PatchDraftRestaurant(c fiber.Ctx, stores *stores.Store, id int) error {
	err := stores.UpdateDraftRestaurant(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"error": err,
			},
		})
	}
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"success": true,
		},
	})
}

func DeleteRestaurant(c fiber.Ctx, stores *stores.Store, id int) error {
	err := stores.DeleteRestaurant(id)

	if err != nil {
		err := c.JSON(fiber.Map{
			"data": fiber.Map{
				"success": false,
				"erreur":  err,
			},
		})
		return err
	}

	err = c.JSON(fiber.Map{
		"data": fiber.Map{
			"success": true,
		},
	})

	return err

}

func CreateNewRestaurantCategory(c fiber.Ctx, stores *stores.Store, item model.RestaurantCategoryItem) error {

	if item.Name == "" {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"error":   "name is empty",
				"success": false,
			},
		})
	}

	verify, err := stores.GetOneCategoryByName(item.Name)

	if verify.Id != 0 {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"error":   "category already exist",
				"success": false,
			},
		})
	}

	res, err := stores.CreateCategory(item)

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
			"id": res,
		},
	})
	return nil
}

func PatchStateRestaurant(c fiber.Ctx, stores *stores.Store, id int) error {
    res, err := stores.GetOneRestaurantById(id)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "data" : fiber.Map{
				"error": err.Error(),
			},
        })
    }

    var newState bool
    if res.Open {
        newState = false
    } else {
        newState = true
    }

    err = stores.UpdateStateRestaurant(id, newState)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "data" : fiber.Map{
				"error": err.Error(),
			},
        })
    }

    return c.JSON(fiber.Map{
		"data" : fiber.Map{
			"state" : newState,
		},
    })
}

