package stores

import (
	"database/sql"
	"foodcourt/app/model"
)

type CommandeStore struct {
	*sql.DB
}

func NewCommandeStore(db *sql.DB) *CommandeStore {
	return &CommandeStore{
		db,
	}
}

func (c *CommandeStore) CreateCommande(model.CommandeItem) (int, error) {
	return 0, nil
}

func (c *CommandeStore) GetCommandeById(id int) (model.CommandeItem, error) {
	return model.CommandeItem{}, nil
}

func (c *CommandeStore) GetAllCommandeByRestaurantId(id int) ([]model.CommandeItem, error) {
	return []model.CommandeItem{}, nil
}
func (c *CommandeStore) UpdateCommande(id int, state int) (model.CommandeItem, error) {
	return model.CommandeItem{}, nil
}
