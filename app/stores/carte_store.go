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

func(c *CarteStore)CreateCarte(carte model.CarteItem) (bool, error) {
	stmt, err := c.Prepare("INSERT INTO carte (restaurant_id, description, price) VALUES (?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(carte.RestaurantId , carte.Description , carte.Price)
	if err != nil {
		return false , err
	}

	return true, nil

}

func (c *CarteStore) DeleteCarteById(carte_id int)(bool, error) {

  _, err := c.Exec("DELETE FROM carte WHERE id = ?", carte_id)

  if err != nil {
	return false, err
  }

  return true,nil
  
}

