package request

import "time"

type CreateCategoryRequestType struct{
	Name string `json:"name"`
}

type CreateCarte struct {
	RestaurantId int `json:"restaurant_id"`
	Description string `json:"description"`
	Price float32 `json:"price"`
}

type CreateProduct struct {
	Produit string `json:"produit"`
	Price float32 `json:"price"`
	CarteId int `json:"carte_id"`
	CategoryId int `json:"category_id"`
}

type CreateCommande struct {
	Date         time.Time `json:"date"`
	UserId       int       `json:"user_id"`
	RestaurantId int       `json:"restaurant_id"`
	Content      []int     `json:"content"`
	Commentaire  string    `json:"commentaire"`
	State        int       `json:"state"`

}

type UpdateCommande struct {
	State        int       `json:"state"`
}