package stores

import (
	"database/sql"
	"foodcourt/app/model"
)

type RolesStore struct {
	*sql.DB
}

func NewRolesStore(db *sql.DB) *RolesStore {
	return &RolesStore{
		db,
	}
}

func (r *RolesStore) GetAllRoles() ([]model.RolesItem, error) {
	return []model.RolesItem{}, nil
}

func (r *RolesStore) GetRoleById(id int) (model.RolesItem, error) {
	return model.RolesItem{}, nil
}
