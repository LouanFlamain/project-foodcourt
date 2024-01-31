package request

type CreateRestaurantRequestType struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"restaurant_name"`
	Description string `json:"description"`
	CategoryId int `json:"category_id"`
}