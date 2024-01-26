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
func(r *RestaurantCategoryStore) CreateCategory(item structure.RestaurantCategoryItem)(int, error){

	res, err := r.DB.Exec("INSERT INTO restaurant_category (name) VALUES (?)", item.Name)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, nil
	}

	return int(id), nil
}
func(r *RestaurantCategoryStore) GetOneCategoryByName(name string)(structure.RestaurantCategoryItem, error){
	var item structure.RestaurantCategoryItem
	rows, err := r.Query("SELECT * FROM restaurant_category WHERE name = ?", name)

	if err != nil {
		return item, err
	}
	defer rows.Close()

	if rows.Next(){
		err := rows.Scan(&item.Id, &item.Name)
		if err != nil{
			return item, err
		}
		return item, nil
	}else{
		return item, sql.ErrNoRows
	}
}
