package main

import (
	"fmt"
	"net/http"
	"os"
	"project/foodcourt/database"
)

func main(){

	//init de la bdd
	db := database.InitDb()
	defer db.Close()

	fmt.Println(os.Getenv("MARIADB_ROOT_PASSWORD"))

	//handler


	//démarre le serveur web

	err := http.ListenAndServe(":8097", nil)

	if err != nil {
		fmt.Printf("le serveur ne peux pas être lancé : %v\n", err)
	}
}