package auth_routes

import (
	"encoding/json"
	"foodcourt/app/api/request"
	"foodcourt/app/api/response"
	"foodcourt/app/handlers"
	"foodcourt/app/middleware"
	"foodcourt/app/stores"
	"log"

	"github.com/gofiber/fiber/v3"
)

func SetUpAuthRoute(route fiber.Router, myStore *stores.Store) {
	route.Post("/register", func(c fiber.Ctx) error {
		var req request.RegisterRequest
		if err := json.Unmarshal(c.Body(), &req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("cannot parse JSON"))
		}

		c.Locals("registerRequest", req)

		responseData, _ := handlers.RegisterHandler(c, myStore, req)

		if emailErr := middleware.RegisterMiddleware(c); emailErr != nil {
			log.Println("Erreur lors de l'envoi de l'email:", emailErr)
		}

		return c.JSON(responseData)
	})

	route.Post("/login", func(c fiber.Ctx) error {
		var req request.LoginRequest
		if err := json.Unmarshal(c.Body(), &req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("cannot parse JSON"))
		}
		return handlers.LoginHandler(c, myStore, req)
	})

	route.Post("/restaurant/create", func(c fiber.Ctx) error {
		var body request.CreateRestaurantRequestType
		err := json.Unmarshal(c.Body(), &body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		return handlers.CreateRestaurant(c, myStore, body)
	})

}
