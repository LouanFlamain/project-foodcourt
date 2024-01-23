package main

import (
	"fmt"
	"net/http"
	"project/foodcourt/database"
)

func main(){

	//init de la bdd
	db := database.InitDb()
	defer db.Close()

	err := http.ListenAndServe(":8097", nil)

	if err != nil {
		fmt.Printf("le serveur ne peux pas être lancé : %v\n", err)
	}
}