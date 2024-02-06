package stores

import (
	"database/sql"
	"foodcourt/app/model"
)

type ProductStore struct {
	*sql.DB
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{
		db,
	}
}

func (p *ProductStore) GetProductsByCarteId(carte_id int) ([]model.ProductItem, error) {
	var productsArray []model.ProductItem

	rows, err := p.Query("SELECT * FROM product WHERE carte_id = ?", carte_id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product model.ProductItem
		err := rows.Scan(&product.Id, &product.Produit, &product.Price, &product.CarteId, &product.CategoryId)

		if err != nil {
			return nil, err
		}

		productsArray = append(productsArray, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productsArray, nil
}
func (p *ProductStore) CreateProduct(product model.ProductItem) (bool, error) {

	stmt, err := p.Prepare("INSERT INTO product (product , price , carte_id , category_id) VALUE (?,?,?,?)")

	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Produit, product.Price, product.CarteId, product.CategoryId)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (p *ProductStore) DeleteProductById(product_id int) (bool, error) {
	_, err := p.Exec("DELETE FROM product WHERE id = ?", product_id)

	if err != nil {
		return false, err
	}

	return true, nil
}
