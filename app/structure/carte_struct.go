package structure


type CarteItem struct {
	Id int `json:"id"`
	RestaurantId int `json:"restaurant_id"`
	Description string `json:"description"`
	Price int `json:"price"`
}



type CarteInterface interface {
  GetMenuByRestaurantId(restaurantId int)(CarteItem,error)
}

