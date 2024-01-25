package store

import (
	"database/sql"
	"project/foodcourt/structure"
)

func CreateStore(db *sql.DB)*Store{
	return &Store{
		UserInterface: NewUserStore(db),
		RestaurantInterface: NewRestaurantStore(db),
		RestaurantCategoryInterface: NewRestaurantCategoryStore(db),
		RolesInterface: NewRolesStore(db),
	}
}

type Store struct{
	structure.UserInterface
	structure.RestaurantCategoryInterface
	structure.RestaurantInterface
	structure.RolesInterface
}