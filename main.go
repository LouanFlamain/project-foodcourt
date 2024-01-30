package main

import (
	"foodcourt/app/config"
	"foodcourt/app/database"
	"foodcourt/app/routes"
	"foodcourt/app/stores"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Chargement de la configuration
	cfg := config.LoadConfig()

	// Initialisation de la base de données avec la configuration chargée
	db := database.InitDb(cfg)

	defer db.Close()

	// Création de l'application Fiber
	app := fiber.New()

	// Instanciation du store avec la base de données
	myStore := stores.CreateStore(db)

	// Configuration des routes
	routes.SetRoute(app, myStore)

	// Récupération du port depuis les variables d'environnement, avec une valeur par défaut
	port := os.Getenv("PORT")
	if port == "" {
		port = "8095" // Port par défaut
	}

	// Démarrage du serveur web sur le port spécifié
	log.Printf("Serveur démarré sur le port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Échec du lancement du serveur : %v", err)
	}
}

type Response struct {
	State string `json:"state"`
	Code  int    `json:"code"`
}
