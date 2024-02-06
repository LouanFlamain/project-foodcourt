package model


type ProductItem struct {
	Id int `json:"id"`
	Produit string `json:"produit"`
	Price float32 `json:"price"`
	CarteId int `json:"carte_id"`
	CategoryId int `json:"category_id"`
}

type CategoryTypeItem struct {
	Id int `json:"id"`
	Name string `json:"name"`
}


type ProductInterface interface {
	GetProductsByCarteId(carte_id int)([]ProductItem, error)
	CreateProduct(ProductItem)(bool, error)
	DeleteProductById(id int)(bool, error)
}


type CategoryInterface interface {
	GetAllCategoryType()([]CategoryTypeItem, error)
	deleteCategoryType(id int)(bool, error)
}

