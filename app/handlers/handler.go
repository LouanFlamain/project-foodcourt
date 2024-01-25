package handlers

import (
	"encoding/json"
	"project/foodcourt/store"

	"github.com/gofiber/fiber/v3"
)

type RequestBody struct {
	Test string `json:"test"`
}

func TestRequest(c fiber.Ctx, store *store.Store) error {
	var body RequestBody

	// Parser le corps de la requête dans la structure 'body'
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		// Gérer l'erreur si le JSON est mal formé
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// (Optionnel) Manipuler ou utiliser 'body' comme nécessaire
	// ...

	// Renvoyer les données parsées (ou modifiées) dans la réponse
	return c.JSON(body) // Ou c.JSON(fiber.Map{"field1": body.Field1, "field2": body.Field2})
}