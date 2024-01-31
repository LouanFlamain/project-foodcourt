package model

type UserItem struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
	Roles    int    `json:"roles"`
}

type UserInterface interface {
	AddUser(item UserItem) (bool, error)
	AddRestaurateur(item UserItem)(int, error)
	GetUsers() ([]UserItem, error)
	GetOneUser(id int) (UserItem, error)
	UpdateUser(user UserItem) (bool, error)
	DeleteUser(id int) (bool, error)
	GetOneUserByUsername(string)(UserItem, error)
	VerifyUserByMail(string)(error)
}

type RolesItem struct {
	Id   int `json:"id"`
	Name int `json:"name"`
}

type RolesInterface interface {
	GetAllRoles() ([]RolesItem, error)
	GetRoleById(id int) (RolesItem, error)
}
