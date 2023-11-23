package roles

import (
	"admin-service-go/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	role := api.Group("/roles")

	role.Post("/create", middleware.Protected(), CreateRole)
	role.Get("/getAllRoles", middleware.Protected(), GetAllRoles)
	role.Get("/:id", middleware.Protected(), GetRoleByRoleId)
	role.Patch("/:id", middleware.Protected(), UpdateRole)
	role.Delete("/:id", middleware.Protected(), DeleteRole)
}
