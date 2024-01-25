package store

import (
	"database/sql"
	"project/foodcourt/structure"
)

type CommandeStore struct{
	*sql.DB
}

func NewCommandeStore(db *sql.DB)*CommandeStore{
	return &CommandeStore{
		db,
	}
}

func(c *CommandeStore) CreateCommande(structure.CommandeItem)(int, error){
	return 0, nil
}

func(c *CommandeStore) GetCommandeById(id int)(structure.CommandeItem, error){
	return structure.CommandeItem{}, nil
}

func(c *CommandeStore) GetAllCommandeByRestaurantId(id int)([]structure.CommandeItem, error){
	return []structure.CommandeItem{}, nil
}
func(c *CommandeStore) UpdateCommande(id int, state int)(structure.CommandeItem, error){
	return structure.CommandeItem{}, nil
}