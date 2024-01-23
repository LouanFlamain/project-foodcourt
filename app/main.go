package main

import (
	"fmt"
	"project/foodcourt/database"
	"project/foodcourt/routes"

	"github.com/gofiber/fiber/v3"
)

func main(){

	//init de la bdd
	db := database.InitDb()
	defer db.Close()

	app := fiber.New();

	//router

	routes.SetRoute(app)

	//démarre le serveur web

	err := app.Listen(":8097")

	if err != nil {
		fmt.Printf("Le lancement du serveur a échoué : %v\n", err)
	}


}

type Response struct {
	State string `json:"state"`
	Code int `json:"code"`
}