package response

import "github.com/gofiber/fiber/v3"

func ErrorResponse(message string) fiber.Map {
	return fiber.Map{"success": false, "error": message}
}

func SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{"success": true, "data": data}
}
