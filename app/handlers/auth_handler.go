package handlers

import (
	"foodcourt/app/api/mercure"
	"foodcourt/app/api/request"
	"foodcourt/app/api/response"
	"foodcourt/app/auth"
	"foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func RegisterHandler(c fiber.Ctx, userStore *stores.Store, req request.RegisterRequest) error {

	existingUser, err := userStore.GetOneUserByEmail(req.Email)
	if err == nil && existingUser.Email != "" {
		return c.JSON(response.ErrorResponse("Email already registered"))
	}

	existingUser, err = userStore.GetOneUserByUsername(req.Username)
	if err == nil && existingUser.Username != "" {
		return c.JSON(response.ErrorResponse("Username already registered"))
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return c.JSON(response.ErrorResponse("Error while hashing password"))
	}

	user := model.UserItem{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	success, err := userStore.AddUser(user)
	if !success || err != nil {
		return c.JSON(response.ErrorResponse(err.Error()))
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		return c.JSON(response.ErrorResponse("cannot generate token"))
	}

	return c.JSON(response.SuccessResponse(fiber.Map{"token": token}))
}

func LoginHandler(c fiber.Ctx, userStore *stores.Store, req request.LoginRequest) error {
	user, err := userStore.GetOneUserByEmail(req.Email)
	if err != nil {
		return c.JSON(response.ErrorResponse("Authentication failed"))
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		return c.JSON(response.ErrorResponse("Authentication failed"))
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		return c.JSON(response.ErrorResponse("Error while generating token"))
	}

	mercure, err := mercure.GenerateJWT()

	if err != nil {
		return c.JSON(response.ErrorResponse("cannot generate mercure token"))
	}

	return c.JSON(response.SuccessResponse(fiber.Map{"message": "login successful", "token": token, "mercure" : mercure}))
}
