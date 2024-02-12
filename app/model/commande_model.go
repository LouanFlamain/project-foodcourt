package model

import (
	// "foodcourt/app/model"
	"time"
)

type CommandeItem struct {
	Id           int           `json:"id"`
	Date         time.Time     `json:"date"`
	UserId       int           `json:"user_id"`
	RestaurantId int           `json:"restaurant_id"`
	Content      []interface{} `json:"content"`
	Commentaire  string        `json:"commentaire"`
	State        int           `json:"state"`
}

type CommandeStateItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CommandeInterface interface {
	CreateCommande(CommandeItem) (bool, error)
	GetCommandeById(id int) (CommandeItem, error)
	GetCommandeByUserId(id int) (CommandeItem, error)
	GetAllCommandeByRestaurantId(id int) ([]CommandeItem, error)
	UpdateCommande(id int, commade CommandeItem) (bool, error)
}
type CommandeStateInterface interface {
	GetAllCommandeState() ([]CommandeStateItem, error)
	GetCommandeStateById(id int) (CommandeStateItem, error)
}
