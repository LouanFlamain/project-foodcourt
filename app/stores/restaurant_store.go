package stores

import (
	"database/sql"
	"foodcourt/app/model"
)

type RestaurantStore struct {
	*sql.DB
}

func NewRestaurantStore(db *sql.DB) *RestaurantStore {
	return &RestaurantStore{
		db,
	}
}

func (r *RestaurantStore) CreateRestaurant(model.RestaurantItem) error {
	return nil
}

func (r *RestaurantStore) UpdateRestaurant(model.RestaurantItem) (model.RestaurantItem, error) {
	return model.RestaurantItem{}, nil
}

func (r *RestaurantStore) GetAllRestaurant() ([]model.RestaurantItem, error) {
	var restaurants []model.RestaurantItem

	rows, err := r.Query("SELECT * FROM restaurant WHERE draft = ?", true)

	if err != nil {
		return []model.RestaurantItem{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurant model.RestaurantItem
		if err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Email, &restaurant.Picture, &restaurant.Description, &restaurant.CategoryId, &restaurant.Draft, &restaurant.Open); err != nil {
			return []model.RestaurantItem{}, err
		}

		restaurants = append(restaurants, restaurant)

	}
	return restaurants, nil
}
func (r *RestaurantStore) GetDraftRestaurant() ([]model.RestaurantItem, error) {
	var restaurants []model.RestaurantItem

	rows, err := r.Query("SELECT * FROM restaurant WHERE draft = ?", false)

	if err != nil {
		return []model.RestaurantItem{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurant model.RestaurantItem
		if err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Email, &restaurant.Picture, &restaurant.Description, &restaurant.CategoryId, &restaurant.Draft, &restaurant.Open); err != nil {
			return []model.RestaurantItem{}, err
		}

		restaurants = append(restaurants, restaurant)

	}
	return restaurants, nil
}

func (r *RestaurantStore) UpdateDraftRestaurant(id int) error {
	_, err := r.Exec("UPDATE restaurant SET draft = true WHERE id = ?", id)

	if err != nil {
		return err
	}
	return err
}

func (r *RestaurantStore) GetAllOpenRestaurant() ([]model.RestaurantItem, error) {
	var restaurants []model.RestaurantItem

	rows, err := r.Query("SELECT id, name, email, picture, description, category_id, open FROM restaurant WHERE open = ?", true)

	if err != nil {
		return []model.RestaurantItem{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurant model.RestaurantItem
		if err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Email, &restaurant.Picture, &restaurant.Description, &restaurant.CategoryId, &restaurant.Open); err != nil {
			return []model.RestaurantItem{}, err
		}
		restaurants = append(restaurants, restaurant)

	}
	return restaurants, nil
}

func (r *RestaurantStore) GetOneRestaurantById(id int) (model.RestaurantItem, error) {
	return model.RestaurantItem{}, nil
}

func (r *RestaurantStore) GetAllRestaurantByCategory(id int) ([]model.RestaurantItem, error) {

	return []model.RestaurantItem{}, nil
}

func (r *RestaurantStore) DeleteRestaurant(id int) error {
	_, err := r.DB.Exec("DELETE FROM restaurant WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *RestaurantStore) UpdateRestaurantOpenState(id int, open bool) (bool, error) {
	return false, nil
}
