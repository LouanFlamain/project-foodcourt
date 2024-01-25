package structure

import "time"

type CommandeItem struct {
	Id int `json:"id"`
	Date time.Time `json:"date"`
	UserId int `json:"user_id"`
	RestaurantId int `json:"restaurant_id"`
	Content string `json:"content"`
	Commentaire string `json:"commentaire"`
	State int `json:"state"`
}

type CommandeStateItem struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type CommandeInterface interface{
	CreateCommande(CommandeItem)(int, error)
	GetCommandeById(id int)(CommandeItem, error)
	GetAllCommandeByRestaurantId(id int)(CommandeItem, error)
	UpdateCommande(id int, state int)(CommandeItem, error)
}
type CommandeStateInterface interface{
	GetAllCommandeState()([]CommandeStateItem, error)
	GetCommandeStateById(id int)(CommandeStateItem, error)
}