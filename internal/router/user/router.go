package user

import "github.com/gofiber/fiber/v2"

func SetupRoutes(api fiber.Router) {
	user := api.Group("/user")

	user.Get("/", CreateUser)
	user.Post("/register", Register)
}
