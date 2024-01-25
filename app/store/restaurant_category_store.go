package store

import (
	"database/sql"
	"project/foodcourt/structure"
)

type RestaurantCategoryStore struct{
	*sql.DB
}

func NewRestaurantCategoryStore(db *sql.DB)*RestaurantCategoryStore{
	return &RestaurantCategoryStore{
		db,
	}
}

func(r *RestaurantCategoryStore) GetOneCategory(id int)(structure.RestaurantCategoryItem, error){
	return structure.RestaurantCategoryItem{}, nil
}

func(r *RestaurantCategoryStore) GetAllCategory()([]structure.RestaurantCategoryItem, error){
	return []structure.RestaurantCategoryItem{}, nil
}
func(r *RestaurantCategoryStore) CreateCategory(structure.RestaurantCategoryItem)(error){
	return nil
}