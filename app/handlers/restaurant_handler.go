package handlers

import (
	"project/foodcourt/store"
	"project/foodcourt/structure"
	"project/foodcourt/web/api/response"

	"github.com/gofiber/fiber/v3"
)

//----------------------------------------------------customer----------------------------------------------------------

func GetAllOpenRestaurant(c fiber.Ctx, store *store.Store) error{
	restaurants, err := store.GetAllOpenRestaurant()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data" : fiber.Map{
				"success" : false,
				"error": "Failed to retrieve restaurants",
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
		"data" : responseRestaurant,
	})
	return err

}
//----------------------------------------------------seller------------------------------------------------------------

//----------------------------------------------------admin-------------------------------------------------------------

func GetAllRestaurant(c fiber.Ctx, store *store.Store) error{
		restaurants, err := store.GetAllRestaurant()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"data" : fiber.Map{
					"success" : false,
					"error": "Failed to retrieve restaurants",
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
			"data" : responseRestaurant,
		})
		return err

}
func GetDraftRestaurant(c fiber.Ctx, store *store.Store) error{
	restaurants, err := store.GetDraftRestaurant()

	if err != nil {
		return c.JSON(fiber.Map{
			"data" : fiber.Map{
				"success" : false,
				"error" : err,
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
			Draft:        restaurant.Draft,
		}
	}

	err = c.JSON(fiber.Map{
		"data" : responseRestaurant,
	})
	return err

}

func PatchDraftRestaurant(c fiber.Ctx, store *store.Store, id int) error {
	err := store.UpdateDraftRestaurant(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"data" : fiber.Map{
				"error" : err,
			},
		})
	}
	return c.JSON(fiber.Map{
		"data" : fiber.Map{
			"success" : true,
		},
	})
}

func DeleteRestaurant(c fiber.Ctx, store *store.Store, id int) error{
	err := store.DeleteRestaurant(id)

	if err != nil {
		err := c.JSON(fiber.Map{
			"data" : fiber.Map{
				"success" : false,
				"erreur" : err,
			},
		})
		return err
	}

	err = c.JSON(fiber.Map{
		"data" : fiber.Map{
			"success" : true,
		},
	})

	return err

}

func CreateNewRestaurantCategory(c fiber.Ctx, store *store.Store, item structure.RestaurantCategoryItem) error{

	
	if item.Name == "" {
		return c.JSON(fiber.Map{
			"data" : fiber.Map{
				"error" : "name is empty",
				"success" : false,
			},
		})
	}

	verify, err := store.GetOneCategoryByName(item.Name)

	if verify.Id != 0 {
		return c.JSON(fiber.Map{
			"data" : fiber.Map{
				"error" : "category already exist",
				"success" : false,
			},
		})
	}


	res, err := store.CreateCategory(item)

	if err != nil {
		 err = c.JSON(fiber.Map{
			"data" : fiber.Map{
				"error" : err,
				"success" : false,
			},
		})
		return err
	}
	err = c.JSON(fiber.Map{
		"data" : fiber.Map{
			"id" : res,
		},
	})
	return nil
}
