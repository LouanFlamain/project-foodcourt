package middleware

import (
	"foodcourt/app/auth"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenString := ""

	if len(authHeader) > 7 && strings.HasPrefix(strings.ToUpper(authHeader), "BEARER ") {
		tokenString = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization header"})
	}

	user, err := auth.ValidateJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: Invalid token"})
	}

	c.Locals("user", user)
	return c.Next()
}
