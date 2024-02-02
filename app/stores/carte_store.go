package stores

import (
	"database/sql"
	"foodcourt/app/model"
)


type CarteStore struct {
	*sql.DB
}

func NewCarteStore(db *sql.DB) *CarteStore {
	return &CarteStore{
		db,
	}
}

func (c *CarteStore) GetCarteByRestaurantId(restaurantId int)(model.CarteItem, error){
	var carte model.CarteItem

	rows, err := c.Query("SELECT * FROM carte WHERE restaurant_id = ? ", restaurantId)

	if err != nil {
		return carte, err
	}

	defer rows.Close()

	if rows.Next() {
		
		err := rows.Scan(&carte.Id , &carte.RestaurantId , &carte.Description , &carte.Price)

		if err != nil {
			return carte , err
		}
		return carte, nil

	} else {
		return carte , sql.ErrNoRows 
	}

}

func (c *CarteStore) DeleteCarteById(carte_id int) error {
  _, err := c.Exec("DELETE FROM CARTE WHERE id = ?", carte_id)

  if err != nil {
	return err
  }

  return nil
  
}

