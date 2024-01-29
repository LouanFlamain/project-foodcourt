package response

type GetDraftRestaurantResponseType struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Picture string `json:"picture"`
	Description string `json:"description"`
	CategoryId int `json:"category_id"`
	Draft bool `json:"draft"`
}
type GetAllRestaurantResponseType struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Picture string `json:"picture"`
	Description string `json:"description"`
	CategoryId int `json:"category_id"`
	Open bool `json:"open"`
}