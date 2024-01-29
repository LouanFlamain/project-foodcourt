package stores

import (
	"database/sql"
	"foodcourt/app/model"
)

type CommandeStateStore struct {
	*sql.DB
}

func NewCommandeStateStore(db *sql.DB) *CommandeStateStore {
	return &CommandeStateStore{
		db,
	}
}

func (c *CommandeStateStore) GetAllCommandeState() ([]model.CommandeStateItem, error) {
	return []model.CommandeStateItem{}, nil
}
func (c *CommandeStateStore) GetCommandeStateById(id int) (model.CommandeStateItem, error) {
	return model.CommandeStateItem{}, nil
}
