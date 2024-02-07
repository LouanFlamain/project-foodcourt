package middleware

import (
	"foodcourt/app/auth"
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
