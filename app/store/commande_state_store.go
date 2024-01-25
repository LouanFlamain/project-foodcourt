package store

import (
	"database/sql"
	"project/foodcourt/structure"
)

type CommandeStateStore struct{
	*sql.DB
}

func NewCommandeStateStore(db *sql.DB)*CommandeStateStore{
	return &CommandeStateStore{
		db,
	}
}

func(c *CommandeStateStore) GetAllCommandeState()([]structure.CommandeStateItem, error){
	return []structure.CommandeStateItem{}, nil
}
func(c *CommandeStateStore) GetCommandeStateById(id int)(structure.CommandeStateItem, error){
	return structure.CommandeStateItem{}, nil
}