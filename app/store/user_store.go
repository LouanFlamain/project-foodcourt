package store

import (
	"database/sql"
	"project/foodcourt/structure"
)

type UserStore struct{
	*sql.DB
}

func NewUserStore(db *sql.DB)*UserStore {
	return &UserStore{
		db,
	}
}

func(u *UserStore) AddUser(structure.UserItem)(bool, error){
	return false, nil
}
func(u *UserStore) GetUsers()([]structure.UserItem, error){
	return []structure.UserItem{}, nil
}
func(u *UserStore) GetOneUser(id int)(structure.UserItem, error){
	return structure.UserItem{}, nil
}
func(u *UserStore) UpdateUser(id int)(bool, error){
	return true, nil
}
func(u *UserStore) DeleteUser(id int) bool{
	return true
}