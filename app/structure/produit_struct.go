package structure


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
	GetAllProducts(carte_id int)([]ProductItem, error)
	DeleteProduct(carte_id int , id int)(error)

}


type CategoryInterface interface {
	GetAllCategoryType()([]CategoryTypeItem, error)
	deleteCategoryType(id int)(error)
}

