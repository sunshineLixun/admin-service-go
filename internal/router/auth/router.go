package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {

	auth := api.Group("/auth")

	auth.Post("/login", Login)
}
