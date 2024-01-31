package model

type FileImageItem struct{
	Name string `json:"name"`
}

type FileImageInterface interface{
	GetImage(id int)(string, error)
}