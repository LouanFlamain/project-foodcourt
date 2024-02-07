package routes

import (
	"foodcourt/app/handlers"
	"foodcourt/app/middleware"
	admin_routes "foodcourt/app/routes/admin"
	auth_routes "foodcourt/app/routes/auth"
	// carte_router "foodcourt/app/routes/carte"
	customer_routes "foodcourt/app/routes/customer"
	seller_routes "foodcourt/app/routes/seller"
	user_routes "foodcourt/app/routes/user"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func SetRoute(app *fiber.App, myStore *stores.Store) {

	//cors

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:3001",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	//création du groupe d'url commençant par /api
	api := app.Group("/api")

	api.Use(func(c fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	//router

	authGroup := api.Group("/auth")
	auth_routes.SetUpAuthRoute(authGroup, myStore)

	userGroup := api.Group("/user", middleware.AuthMiddleware)
	user_routes.SetUpUserRoute(userGroup, myStore)

	adminGroup := api.Group("/admin", middleware.CheckAdminMiddleware(myStore, 3))
	admin_routes.SetUpAdminRoute(adminGroup, myStore)

	customerGroup := api.Group("/customer", middleware.AuthMiddleware)
	customer_routes.SetUpCustomerRoute(customerGroup, myStore)

	sellerGroup := api.Group("/seller", middleware.AuthMiddleware, middleware.CheckAdminMiddleware(myStore, 2))
	seller_routes.SetUpSellerRoute(sellerGroup, myStore)

	api.Get("/test", func(c fiber.Ctx) error {
		response := Response{true, "/api/test"}

		return c.JSON(response)
	})

	type MyRequestBody struct {
		Test string `json:"test"`
	}
	app.Post("/test", func(c fiber.Ctx) error {
		return handlers.GetAllRestaurant(c, myStore)
	})
}

type Response struct {
	Success  bool   `json:"success"`
	Location string `json:"location"`
}
