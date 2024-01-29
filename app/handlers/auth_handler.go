package handlers

import (
	"encoding/json"
	"foodcourt/app/api/response"
	"foodcourt/app/auth"
	"foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func RegisterHandler(c fiber.Ctx, userStore *stores.Store) error {
	var user model.UserItem
	//mettre structure request
	err := json.Unmarshal(c.Body(), "la tu mets la struct de la request")
	if err != nil {
		return c.JSON(fiber.Map{
			"data": response.ErrorResponse(err.Error()),
		})
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": response.ErrorResponse("Error while hashing password"),
		})
	}
	user.Password = hashedPassword

	success, err := userStore.AddUser(user)
	if !success || err != nil {
		return c.JSON(fiber.Map{
			"data": response.ErrorResponse(err.Error()),
		})
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": response.ErrorResponse("cannot generate token"),
		})
	}

	return c.JSON(fiber.Map{
		"data": response.SuccessResponse(fiber.Map{"token": token}),
	})
}

func LoginHandler(c fiber.Ctx, userStore *stores.Store) error {
	var credentials model.UserItem
	//metttre la structure de la requête
	err := json.Unmarshal(c.Body(), "la tu mets la struct de la request")
	if err != nil {
		return c.JSON(fiber.Map{
			"data": response.ErrorResponse(err.Error()),
		})
	}

	user, err := userStore.GetOneUserByUsername(credentials.Username)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": response.ErrorResponse("Authentication failed"),
		})
	}

	if !auth.CheckPasswordHash(credentials.Password, user.Password) {
		return c.JSON(fiber.Map{
			"data": response.ErrorResponse("Authentication failed"),
		})
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": response.ErrorResponse("Error while generating token"),
		})
	}

	return c.JSON(fiber.Map{
		"data": response.SuccessResponse(fiber.Map{"message": "login successful", "token": token}),
	})
}
