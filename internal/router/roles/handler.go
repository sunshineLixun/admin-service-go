package roles

import "github.com/gofiber/fiber/v2"

func GetAllRoles(c *fiber.Ctx) error {
	return c.SendString("GetAllRoles")
}
