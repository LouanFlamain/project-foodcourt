package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"foodcourt/app/auth"
	"foodcourt/app/stores"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	var tokenString string

	if len(authHeader) > 7 && strings.HasPrefix(strings.ToUpper(authHeader), "BEARER ") {
		tokenString = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization header"})
	}

	userClaims, err := auth.ValidateJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: Invalid token"})
	}

	newUserClaims := &auth.UserClaims{
		Id:       userClaims.Id,
		Username: userClaims.Username,
		Email:    userClaims.Email,
	}

	c.Locals("userClaims", newUserClaims)
	return c.Next()
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

type User struct {
	Id int `json:"id"`
}

type PayloadParse struct {
	User User `json:"user"`
	Exp  int  `json:"exp"`
}

func CheckAdminMiddleware(store *stores.Store, roleId int) fiber.Handler {
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

		user, err := store.GetOneUser(payloadParser.User.Id)

		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(fiber.Map{
				"data": fiber.Map{
					"success": false,
					"error":   err,
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
