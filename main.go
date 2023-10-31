package main

import (
	"admin-service-go/database"
	"admin-service-go/internal/router"
	"log"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "admin-service-go/docs"
)

func main() {

	if err := database.ConnectDb(); err != nil {
		log.Panic("链接失败, 错误原因", err.Error())
	}

	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":8080"))

}
