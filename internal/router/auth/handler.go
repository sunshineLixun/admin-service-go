package auth

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello", "data": nil})
}

func Login(c *fiber.Ctx) error {
	// TODO: login
	return nil
}
