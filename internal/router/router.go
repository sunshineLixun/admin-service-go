package router

import (
	"admin-service-go/internal/router/auth"
	"admin-service-go/internal/router/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	auth.SetupRoutes(api)

	user.SetupRoutes(api)

}
