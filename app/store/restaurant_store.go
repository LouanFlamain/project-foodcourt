package store

import (
	"database/sql"
	"project/foodcourt/structure"
)

type RestaurantStore struct{
	*sql.DB
}

func NewRestaurantStore(db *sql.DB)*RestaurantStore{
	return &RestaurantStore{
		db,
	}
}

func(r *RestaurantStore) CreateRestaurant(structure.RestaurantItem)error{
	return nil
}

func(r *RestaurantStore) UpdateRestaurant(structure.RestaurantItem)(structure.RestaurantItem, error){
	return structure.RestaurantItem{}, nil
}
func(r *RestaurantStore) GetAllRestaurant()([]structure.RestaurantItem, error){
	return []structure.RestaurantItem{}, nil
}
func(r *RestaurantStore) GetOneRestaurant(id int)(structure.RestaurantItem, error){
	return structure.RestaurantItem{}, nil
}
func(r *RestaurantStore) GetAllRestaurantByCategory(id int)([]structure.RestaurantItem, error){
	return []structure.RestaurantItem{}, nil
}
func(r *RestaurantStore) DeleteRestaurant(id int)(error){
	return nil
}