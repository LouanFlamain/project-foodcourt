package handlers

import (
	"fmt"
	"foodcourt/app/api/request"
	"foodcourt/app/api/response"
	"foodcourt/app/auth"
	"foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

func GetMyUser(c fiber.Ctx, userStore *stores.Store) error {
	userClaimsInterface := c.Locals("userClaims")
	userClaims, ok := userClaimsInterface.(*auth.UserClaims)
	if !ok || userClaims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	user, err := userStore.GetOneUser(userClaims.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
	}

	user.Password = ""
	return c.JSON(user)
}

func GetUserByID(c fiber.Ctx, userStore *stores.Store) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("Invalid user ID"))
	}

	user, err := userStore.GetOneUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse("User not found"))
	}

	return c.JSON(response.SuccessResponse(user))
}

func GetAllUsers(c fiber.Ctx, userStore *stores.Store) error {
	users, err := userStore.GetUsers()
	if err != nil {
		fmt.Printf("Error retrieving users: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse("Failed to retrieve users"))
	}

	fmt.Printf("Retrieved %d users\n", len(users))
	return c.JSON(response.SuccessResponse(users))
}

func UpdateUser(c fiber.Ctx, userStore *stores.Store, userUpdates request.UpdateUserRequest) error {
	userClaimsInterface := c.Locals("userClaims")
	userClaims, ok := userClaimsInterface.(*auth.UserClaims)
	if !ok || userClaims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	success, err := userStore.UpdateUser(model.UserItem{
		Id:       userClaims.Id,
		Username: userUpdates.Username,
		Email:    userUpdates.Email,
		Picture:  userUpdates.Picture,
	})

	if err != nil || !success {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse("Failed to update user"))
	}

	return c.JSON(response.SuccessResponse(fiber.Map{"message": "User updated successfully"}))
}

func DeleteUser(c fiber.Ctx, userStore *stores.Store) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("Invalid user ID"))
	}

	success, err := userStore.DeleteUser(userID)
	if err != nil || !success {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse("Failed to delete user"))
	}

	return c.JSON(response.SuccessResponse(fiber.Map{"message": "User deleted successfully"}))
}

func ChangePassword(c fiber.Ctx, userStore *stores.Store, passwordChangeRequest request.PasswordChangeRequest) error {
	userClaimsInterface := c.Locals("userClaims")
	userClaims, ok := userClaimsInterface.(*auth.UserClaims)
	if !ok || userClaims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	user, err := userStore.GetOneUser(userClaims.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse("User not found"))
	}

	if !auth.CheckPasswordHash(passwordChangeRequest.OldPassword, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("Incorrect old password"))
	}

	hashedNewPassword, err := auth.HashPassword(passwordChangeRequest.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse("Failed to hash new password"))
	}

	success, err := userStore.UpdateUserPassword(userClaims.Id, hashedNewPassword)
	if !success || err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse("Failed to update password"))
	}

	return c.JSON(response.SuccessResponse(fiber.Map{"message": "Password changed successfully"}))
}
