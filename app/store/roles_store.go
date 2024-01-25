package store

import (
	"database/sql"
	"project/foodcourt/structure"
)

type RolesStore struct{
	*sql.DB
}

func NewRolesStore(db *sql.DB)*RolesStore{
	return &RolesStore{
		db,
	}
}

func(r *RolesStore) GetAllRoles()([]structure.RolesItem, error){
	return []structure.RolesItem{}, nil
}

func(r *RolesStore) GetRoleById(id int)(structure.RolesItem, error){
	return structure.RolesItem{}, nil
}