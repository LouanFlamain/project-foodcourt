package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"foodcourt/app/stores"
	"strings"

	"github.com/gofiber/fiber/v3"
)


type PayloadParse struct {
	Id int `json:"id"`
}

func decodeBase64UrlSafe(encoded string) (string, error) {
	// Remplacer les caractères URL-safe par ceux de l'encodage Base64 standard
	encoded = strings.ReplaceAll(encoded, "-", "+")
	encoded = strings.ReplaceAll(encoded, "_", "/")

	// Ajouter le padding manquant si nécessaire
	switch len(encoded) % 4 {
	case 2:
		encoded += "=="
	case 3:
		encoded += "="
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(decodedBytes), nil
}

func splitJWT(tokenString string) (string, string, string, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("invalid JWT format")
	}

	header, err := decodeBase64UrlSafe(parts[0])
	if err != nil {
		return "", "", "", err
	}

	payload, err := decodeBase64UrlSafe(parts[1])
	if err != nil {
		return "", "", "", err
	}

	// Pas besoin de décoder la signature pour la plupart des utilisations
	signature := parts[2]

	return header, payload, signature, nil
}

func CheckPermissionMiddleware(store *stores.Store, roleId int) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		var tokenString string
		var payloadParser PayloadParse

		if len(authHeader) > 7 && strings.HasPrefix(strings.ToUpper(authHeader), "BEARER ") {
			tokenString = authHeader[7:]
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization header"})
		}

		_, payload, _, err := splitJWT(tokenString)
		if err != nil {
			fmt.Println("Erreur :", err)
			return err
		}

		err = json.Unmarshal([]byte(payload), &payloadParser)
		if err != nil {
			fmt.Println("Erreur lors du décodage du payload:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
		}

		user, err := store.GetOneUser(payloadParser.Id)

		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(fiber.Map{
				"data": fiber.Map{
					"success": false,
					"error":   err,
					"fail": true,
				},
			})
		}

		if user.Roles != roleId {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"data": fiber.Map{
					"success": false,
					"error":   "don't have permission",
				},
			})
		}
		return c.Next()
	}
}
