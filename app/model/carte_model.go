package model


type CarteItem struct {
	Id int `json:"id"`
	RestaurantId int `json:"restaurant_id"`
	Description string `json:"description"`
	Price float32 `json:"price"`
}

type CarteInterface interface {
  GetCarteByRestaurantId(restaurantId int)(CarteItem,error)
  CreateCarte(CarteItem)(bool, error)
  DeleteCarteById(carte_id int)(bool, error)
}

