package structure

type RestaurantItem struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Picture string `json:"picture"`
	Description string `json:"description"`
	CategoryId string `json:"category_id"`
	Draft bool `json:"draft"`
	Open bool `json:"open"`
}

type RestaurantCategoryItem struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

type RestaurantInterface interface{
	CreateRestaurant(RestaurantItem)(error)
	UpdateRestaurant(RestaurantItem)(RestaurantItem, error)
	GetAllRestaurant()([]RestaurantItem, error)
	GetOneRestaurantById(id int)(RestaurantItem, error)
	GetAllRestaurantByCategory(category_id int)([]RestaurantItem, error)
	DeleteRestaurant(id int)(error)
	UpdateRestaurantOpenState(id int, open bool)(bool, error)
}

type RestaurantCategoryInterface interface{
	GetOneCategory(id int)(RestaurantCategoryItem, error)
	GetAllCategory()([]RestaurantCategoryItem, error)
	CreateCategory(RestaurantCategoryItem)(error)
}