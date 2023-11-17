package roles

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	user := api.Group("/roles")

	user.Get("/", GetAllRoles)
}
