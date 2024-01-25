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
	var restaurants []structure.RestaurantItem

	rows, err := r.Query("SELECT * FROM restaurant")

	if err != nil {
		return []structure.RestaurantItem{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurant  structure.RestaurantItem
		if err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Email, &restaurant.Picture, &restaurant.Description, &restaurant.CategoryId, &restaurant.Draft, &restaurant.Open ); err != nil {
			return []structure.RestaurantItem{}, err
		}

		restaurants = append(restaurants, restaurant)
		
	}
	return restaurants, nil
}


func(r *RestaurantStore) GetOneRestaurantById(id int)(structure.RestaurantItem, error){
	return structure.RestaurantItem{}, nil
}


func(r *RestaurantStore) GetAllRestaurantByCategory(id int)([]structure.RestaurantItem, error){
	var restaurants []structure.RestaurantItem

	rows, err := r.Query("SELECT * FROM restaurant")

	if err != nil {
		return []structure.RestaurantItem{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurant  structure.RestaurantItem
		if err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Email, &restaurant.Picture, &restaurant.Description, &restaurant.CategoryId, &restaurant.Draft, &restaurant.Open ); err != nil {
			return []structure.RestaurantItem{}, err
		}

		restaurants = append(restaurants, restaurant)
		
	}
	return restaurants, nil
}


func(r *RestaurantStore) DeleteRestaurant(id int)(error){
	_, err := r.DB.Exec("DELETE FROM restaurant WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}


func(r *RestaurantStore) UpdateRestaurantOpenState(id int, open bool)(bool, error){
	return false, nil
}