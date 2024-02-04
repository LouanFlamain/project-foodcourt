package user_routes

import (
	"encoding/json"
	"foodcourt/app/api/request"
	"foodcourt/app/handlers"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func SetUpUserRoute(route fiber.Router, myStore *stores.Store) {
	route.Get("other/:id", func(c fiber.Ctx) error {
		return handlers.GetUserByID(c, myStore)
	})

	route.Get("/all", func(c fiber.Ctx) error {
		return handlers.GetAllUsers(c, myStore)
	})

	route.Get("/", func(c fiber.Ctx) error {
		return handlers.GetMyUser(c, myStore)
	})

	route.Put("/", func(c fiber.Ctx) error {
		var userUpdates request.UpdateUserRequest
		if err := json.Unmarshal(c.Body(), &userUpdates); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
		}
		return handlers.UpdateUser(c, myStore, userUpdates)
	})

	route.Put("/password", func(c fiber.Ctx) error {
		var passwordChange request.PasswordChangeRequest
		if err := json.Unmarshal(c.Body(), &passwordChange); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
		}
		return handlers.ChangePassword(c, myStore, passwordChange)
	})

	route.Delete("/:id", func(c fiber.Ctx) error {
		return handlers.DeleteUser(c, myStore)
	})
}
