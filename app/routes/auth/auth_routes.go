package auth_routes

import (
	"encoding/json"
	"foodcourt/app/api/request"
	"foodcourt/app/api/response"
	"foodcourt/app/handlers"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func SetUpAuthRoute(route fiber.Router, myStore *stores.Store) {
	route.Post("/register", func(c fiber.Ctx) error {
		var req request.RegisterRequest
		if err := json.Unmarshal(c.Body(), &req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("cannot parse JSON"))
		}
		return handlers.RegisterHandler(c, myStore, req)
	})

	route.Post("/login", func(c fiber.Ctx) error {
		var req request.LoginRequest
		if err := json.Unmarshal(c.Body(), &req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("cannot parse JSON"))
		}
		return handlers.LoginHandler(c, myStore, req)
	})
}
