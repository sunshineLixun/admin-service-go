package user

import "github.com/gofiber/fiber/v2"

func GetUser() {

}

func CreateUser(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"status": "success", "message": "Created user", "data": nil})
}

func UpdateUser() {

}

func DeleteUser() {

}
