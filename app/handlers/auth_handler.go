package handlers

import (
	"encoding/json"
	"foodcourt/app/api/request"
	"foodcourt/app/api/response"
	"foodcourt/app/auth"
	"foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func RegisterHandler(c fiber.Ctx, userStore *stores.Store) error {
	var req request.RegisterRequest
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return c.JSON(fiber.Map{"data": response.ErrorResponse(err.Error())})
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return c.JSON(fiber.Map{"data": response.ErrorResponse("Error while hashing password")})
	}

	user := model.UserItem{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	success, err := userStore.AddUser(user)
	if !success || err != nil {
		return c.JSON(fiber.Map{"data": response.ErrorResponse(err.Error())})
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		return c.JSON(fiber.Map{"data": response.ErrorResponse("cannot generate token")})
	}

	return c.JSON(fiber.Map{"data": response.SuccessResponse(fiber.Map{"token": token})})
}

func LoginHandler(c fiber.Ctx, userStore *stores.Store) error {
	var req request.LoginRequest
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return c.JSON(fiber.Map{"data": response.ErrorResponse(err.Error())})
	}

	user, err := userStore.GetOneUserByUsername(req.Email)
	if err != nil {
		return c.JSON(fiber.Map{"data": response.ErrorResponse("Authentication failed")})
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		return c.JSON(fiber.Map{"data": response.ErrorResponse("Authentication failed")})
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		return c.JSON(fiber.Map{"data": response.ErrorResponse("Error while generating token")})
	}

	return c.JSON(fiber.Map{"data": response.SuccessResponse(fiber.Map{"message": "login successful", "token": token})})
}
