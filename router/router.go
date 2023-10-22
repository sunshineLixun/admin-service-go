package router

import (
	"admin-service-go/router/auth"
	"admin-service-go/router/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	auth.SetupRoutes(app)

	user.SetupRoutes(app)

}
