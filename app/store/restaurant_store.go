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

	rows, err := r.Query("SELECT * FROM restaurant WHERE draft = ?", true)

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
func(r *RestaurantStore) GetDraftRestaurant()([]structure.RestaurantItem, error){
	var restaurants []structure.RestaurantItem

	rows, err := r.Query("SELECT * FROM restaurant WHERE draft = ?", false)

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

func(r *RestaurantStore) UpdateDraftRestaurant(id int)(error){
	_, err := r.Exec("UPDATE restaurant SET draft = true WHERE id = ?", id)

	if err != nil {
		return err
	}
	return err
}

func(r *RestaurantStore) GetAllOpenRestaurant()([]structure.RestaurantItem, error){
	var restaurants []structure.RestaurantItem

	rows, err := r.Query("SELECT id, name, email, picture, description, category_id, open FROM restaurant WHERE open = ?", true)

	if err != nil {
		return []structure.RestaurantItem{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurant  structure.RestaurantItem
		if err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Email, &restaurant.Picture, &restaurant.Description, &restaurant.CategoryId, &restaurant.Open ); err != nil {
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

	return []structure.RestaurantItem{}, nil
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