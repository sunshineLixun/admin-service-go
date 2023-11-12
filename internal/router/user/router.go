package user

import (
	"admin-service-go/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	user := api.Group("/user")

	user.Get("/", GetAllUser)
	user.Post("/register", Register)
	user.Get("/:id", GetUserById)
	user.Delete("/:id", middleware.Protected(), DeleteUser)
	user.Patch("/:id", UpdateUser)
}
