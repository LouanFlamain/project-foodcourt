package routes

import (
	admin_routes "project/foodcourt/routes/admin"
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


	//router vers admin (en suivant api)

	adminGroup := api.Group("/admin")
	admin_routes.SetUpAdminRoute(adminGroup)

	//router vers customer (en suivant api)

	customerGroup := api.Group("/customer")
	customer_routes.SetUpCustomerRoute(customerGroup)

	//router vers seller (en suivant api)

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
