package stores

import (
	"database/sql"
	"foodcourt/app/model"
)

func CreateStore(db *sql.DB) *Store {
	return &Store{
		UserInterface:               NewUserStore(db),
		RestaurantInterface:         NewRestaurantStore(db),
		RestaurantCategoryInterface: NewRestaurantCategoryStore(db),
		RolesInterface:              NewRolesStore(db),
		CarteInterface:              NewCarteStore(db),
		ProductInterface:            NewProductStore(db),
		CommandeInterface:           NewCommandeStore(db),
		
	}
}

type Store struct {
	model.UserInterface
	model.RestaurantCategoryInterface
	model.RestaurantInterface
	model.RolesInterface
	model.CarteInterface
	model.ProductInterface
	model.CommandeInterface
}
