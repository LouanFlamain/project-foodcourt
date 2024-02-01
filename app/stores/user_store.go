package stores

import (
	"database/sql"
	"foodcourt/app/model"
)

type UserStore struct {
	*sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

func (u *UserStore) AddUser(user model.UserItem) (bool, error) {
	stmt, err := u.Prepare("INSERT INTO users (username, email, password, picture, roles) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.Picture, user.Roles)
	if err != nil {
		return false, err
	}

	return true, nil
}
func (u *UserStore) AddRestaurateur(user model.UserItem) (int, error) {
	res, err := u.DB.Exec("INSERT INTO users (username, password, email, roles) VALUES (?, ?, ?, ?)", user.Username, user.Password, user.Email, user.Roles)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (u *UserStore) GetUsers() ([]model.UserItem, error) {
	rows, err := u.Query("SELECT id, username, email, password, picture, roles FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.UserItem
	for rows.Next() {
		var user model.UserItem
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Picture, &user.Roles); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserStore) GetOneUser(id int) (model.UserItem, error) {
	var user model.UserItem
	err := u.QueryRow("SELECT id, username, email, password, picture, roles FROM users WHERE id = ?", id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Picture, &user.Roles)
	if err != nil {
		return model.UserItem{}, err
	}

	return user, nil
}
func (u *UserStore) VerifyUserByMail(email string) error {
	var user model.UserItem
	err := u.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Picture, &user.Roles)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserStore) UpdateUser(user model.UserItem) (bool, error) {
	stmt, err := u.Prepare("UPDATE users SET username = ?, email = ?, password = ?, picture = ?, roles = ? WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.Picture, user.Roles, user.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserStore) DeleteUser(id int) (bool, error) {
	stmt, err := u.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserStore) GetOneUserByUsername(username string) (model.UserItem, error) {
	var user model.UserItem
	err := u.QueryRow("SELECT id, username, email, password, picture, roles FROM users WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Picture, &user.Roles)
	if err != nil {
		return model.UserItem{}, err
	}

	return user, nil
}

func (u *UserStore) GetOneUserByEmail(email string) (model.UserItem, error) {
	var user model.UserItem
	err := u.QueryRow("SELECT id, username, email, password, picture, roles FROM users WHERE email = ?", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Picture, &user.Roles)
	if err != nil {
		return model.UserItem{}, err
	}

	return user, nil
}
