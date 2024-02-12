package stores

import (
	"database/sql" //////////////////////////////////date , user_id , restaurant_id , content , commantaire , stat
	"encoding/json"
	"fmt"
	"foodcourt/app/model"
	"time"
)

type CommandeStore struct {
	*sql.DB
}

func NewCommandeStore(db *sql.DB) *CommandeStore {
	return &CommandeStore{
		db,
	}
}

func (c *CommandeStore) CreateCommande(commande model.CommandeItem) (bool, error) {

	stmt, err := c.Prepare("INSERT INTO commande (date , user_id , restaurant_id , content , commentaire , state ) VALUE (?,?,?,?,?,?)")

	if err != nil {
		return false, err
	}
	defer stmt.Close()

	contentJSON, err := json.Marshal(commande.Content)
	fmt.Println("je suis le casse couille ", &contentJSON)

	if err != nil {
		fmt.Println("Erreur lors de la conversion du contenu en JSON:", err)
		return false, err
	}

	_, err = stmt.Exec(commande.Date, commande.UserId, commande.RestaurantId, contentJSON, commande.Commentaire, commande.State)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *CommandeStore) GetCommandeById(id int) (model.CommandeItem, error) {
	var commande model.CommandeItem

	rows, err := c.Query("SELECT * FROM commande WHERE id = ?", id)
	if err != nil {
		return commande, err
	}
	defer rows.Close()

	if rows.Next() {
		var dateStr string
		var contentJSON string
		err := rows.Scan(&commande.Id, &dateStr, &commande.UserId, &commande.RestaurantId, &contentJSON, &commande.Commentaire, &commande.State)
		if err != nil {
			return commande, err
		}

		// Convertir la chaîne de date en time.Time
		commande.Date, err = time.Parse("2006-01-02 15:04:05", dateStr)
		if err != nil {
			return commande, err
		}

		// Décodez le contenu JSON dans la structure de données Go appropriée
		var content []interface{}
		err = json.Unmarshal([]byte(contentJSON), &content)
		if err != nil {
			return commande, err
		}
		commande.Content = content

		return commande, nil
	} else {
		return commande, sql.ErrNoRows
	}
}

func (c *CommandeStore) GetAllCommandeByRestaurantId(id int) ([]model.CommandeItem, error) {
	var commandesArray []model.CommandeItem

	rows, err := c.Query("SELECT * FROM commande WHERE restaurant_id = ?", id)

	if err != nil {
		return []model.CommandeItem{}, err
	}

	defer rows.Close()

	for rows.Next() {

		var commande model.CommandeItem
		var dateStr string
		var contentJSON []byte
		err := rows.Scan(&commande.Id, &dateStr, &commande.UserId, &commande.RestaurantId, &contentJSON, &commande.Commentaire, &commande.State)

		if err != nil {
			return []model.CommandeItem{}, err
		}
		// Convertir la chaîne de date en time.Time
		commande.Date, err = time.Parse("2006-01-02 15:04:05", dateStr)
		if err != nil {
			return []model.CommandeItem{}, err
		}

		// Convertir le JSON en tableau d'entiers
		var content []interface{}
		if err := json.Unmarshal(contentJSON, &content); err != nil {
			return nil, err
		}
		commande.Content = content

		commandesArray = append(commandesArray, commande)
	}

	if err := rows.Err(); err != nil {
		return []model.CommandeItem{}, err
	}

	return commandesArray, nil
}
func (c *CommandeStore) UpdateCommande(id int, commade model.CommandeItem) (bool, error) {

	stmt, err := c.Prepare("UPDATE commande SET state = ? WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(commade.State, id)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *CommandeStore) GetCommandeByUserId(id int) (model.CommandeItem, error) {
	var commande model.CommandeItem

	rows, err := c.Query("SELECT * FROM commande WHERE user_id = ?", id)
	if err != nil {
		return commande, err
	}
	defer rows.Close()

	if rows.Next() {
		var dateStr string
		var contentJSON string
		err := rows.Scan(&commande.Id, &dateStr, &commande.UserId, &commande.RestaurantId, &contentJSON, &commande.Commentaire, &commande.State)
		if err != nil {
			return commande, err
		}

		// Convertir la chaîne de date en time.Time
		commande.Date, err = time.Parse("2006-01-02 15:04:05", dateStr)
		if err != nil {
			return commande, err
		}

		// Décodez le contenu JSON dans la structure de données Go appropriée
		var content []interface{}
		err = json.Unmarshal([]byte(contentJSON), &content)
		if err != nil {
			return commande, err
		}
		commande.Content = content

		return commande, nil
	} else {
		return commande, sql.ErrNoRows
	}
}
