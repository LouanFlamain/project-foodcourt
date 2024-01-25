package structure

type UserItem struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Picture string `json:"picture"`
	Roles int `json:"roles"`
}

type UserInterface interface{
	AddUser(item UserItem)(bool, error)
	GetUsers()([]UserItem, error)
	GetOneUser(id int)(UserItem, error)
	UpdateUser(id int)(bool, error)
	DeleteUser(id int)(bool)
}

type RolesItem struct {
	Id int `json:"id"`
	Name int `json:"name"`
}

type RolesInterface interface {
	GetAllRoles()([]RolesItem, error)
	GetRoleById(id int)(RolesItem, error)
}