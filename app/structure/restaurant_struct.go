package structure

type RestaurantItem struct{
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Picture string `json:"picture"`
	Location string `json:"location"`
	Description string `json:"description"`
	CategoryId string `json:"category_id"`
	Draft bool `json:"draft"`
}

type RestaurantCategoryItem struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

type RestaurantInterface interface{
	CreateRestaurant(RestaurantItem)(error)
	UpdateRestaurant(RestaurantItem)(RestaurantItem, error)
	GetAllRestaurant()([]RestaurantCategoryItem, error)
	GetOneRestaurant(id int)(RestaurantCategoryItem, error)
	GetAllRestaurantByCategory(category_id int)([]RestaurantCategoryItem, error)
	DeleteRestaurant(id int)(error)
}

type RestaurantCategoryInterface interface{
	GetOneCategory(id int)(RestaurantCategoryItem, error)
	GetAllCategory()([]RestaurantCategoryItem, error)
	CreateCategory(RestaurantCategoryItem)(error)
}