package main

import (
	"admin-service-go/global"
	"admin-service-go/internal/models"
	"admin-service-go/internal/router"
	"fmt"
	"log"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "admin-service-go/docs"
)

func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("读取配置初始化错误:%v", err)
	}

	err = models.SetupDBEngine()

	if err != nil {
		log.Fatalf("配置数据库错误:%v", err)
	}

}

func main() {

	app := fiber.New(fiber.Config{
		ReadTimeout:  global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
	})

	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", global.ServerSetting.HttpHost, global.ServerSetting.HttpPort)))

}
