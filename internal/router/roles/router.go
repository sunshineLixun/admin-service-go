package roles

import (
	"admin-service-go/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	role := api.Group("/roles")

	role.Post("/create", middleware.Protected(), CreateRole)
	role.Get("/", middleware.Protected(), GetAllRoles)
}
