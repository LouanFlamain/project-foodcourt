package routes

import (
	admin_routes "project/foodcourt/routes/admin"
	auth_routes "project/foodcourt/routes/auth"
	customer_routes "project/foodcourt/routes/customer"
	seller_routes "project/foodcourt/routes/seller"

	"github.com/gofiber/fiber/v3"
)

func SetRoute(app *fiber.App){

	//création du groupe d'url commençant par /api
	api := app.Group("/api")    

	api.Use(func(c fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	//router

	authGroup := api.Group("/auth")
	auth_routes.SetUpAuthRoute(authGroup)

	adminGroup := api.Group("/admin")
	admin_routes.SetUpAdminRoute(adminGroup)

	customerGroup := api.Group("/customer")
	customer_routes.SetUpCustomerRoute(customerGroup)

	sellerGroup := api.Group("/seller")
	seller_routes.SetUpSellerRoute(sellerGroup)



	api.Get("/test", func(c fiber.Ctx) error {
		response := Response{true, "/api/test"}

		return c.JSON(response)
	})
	app.Get("/test", func(c fiber.Ctx) error {
		response := Response{true, "/test"}

		return c.JSON(response)
	})
}

type Response struct {
	Success bool `json:"success"`
	Location string `json:"location"`
}
