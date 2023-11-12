package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {

	auth := api.Group("/auth")

	auth.Get("/", Hello)

	auth.Post("/login", Login)
}
