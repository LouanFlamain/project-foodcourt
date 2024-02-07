package handlers

import (
	"foodcourt/app/api/request"
	"foodcourt/app/api/response"
	"foodcourt/app/auth"
	"foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func RegisterHandler(c fiber.Ctx, userStore *stores.Store, req request.RegisterRequest) (*fiber.Map, error) {
	var errorResponse fiber.Map

	existingUser, err := userStore.GetOneUserByEmail(req.Email)
	if err == nil && existingUser.Email != "" {
		errorResponse = response.ErrorResponse("Email already registered")
		return &errorResponse, nil
	}

	existingUser, err = userStore.GetOneUserByUsername(req.Username)
	if err == nil && existingUser.Username != "" {
		errorResponse = response.ErrorResponse("Username already registered")
		return &errorResponse, nil
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		errorResponse = response.ErrorResponse("Error while hashing password")
		return &errorResponse, nil
	}

	user := model.UserItem{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	success, err := userStore.AddUser(user)
	if !success || err != nil {
		errorResponse = response.ErrorResponse(err.Error())
		return &errorResponse, nil
	}

	token, err := auth.GenerateJWT(user.Id, user.Username, user.Email)
	if err != nil {
		errorResponse = response.ErrorResponse("Cannot generate token")
		return &errorResponse, nil
	}

	responseData := response.SuccessResponse(fiber.Map{"token": token})
	return &responseData, nil
}

func LoginHandler(c fiber.Ctx, userStore *stores.Store, req request.LoginRequest) error {
	user, err := userStore.GetOneUserByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("Authentication failed: User not found"))
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("Authentication failed: Incorrect password"))
	}

	token, err := auth.GenerateJWT(user.Id, user.Username, user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse("Error while generating token"))
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse(fiber.Map{"message": "Login successful", "token": token}))
}
