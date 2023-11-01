package router

import (
	"admin-service-go/internal/router/auth"
	"admin-service-go/internal/router/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1", logger.New())

	auth.SetupRoutes(api)

	user.SetupRoutes(api)

}
