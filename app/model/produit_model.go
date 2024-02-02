package model


type ProductItem struct {
	Id int `json:"id"`
	Produit string `json:"produit"`
	Price int `json:"price"`
	CarteId int `json:"carte_id"`
	CategoryId int `json:"category_id"`
}

type CategoryTypeItem struct {
	Id int `json:"id"`
	Name string `json:"name"`
}


type ProductInterface interface {
	GetProductsByCarteId(carte_id int)([]ProductItem, error)
	DeleteProductById(id int)(error)
}


type CategoryInterface interface {
	GetAllCategoryType()([]CategoryTypeItem, error)
	deleteCategoryType(id int)(error)
}

