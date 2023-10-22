package user

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Get("/", CreateUser)
}
