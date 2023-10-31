package user

import (
	"admin-service-go/database"
	"admin-service-go/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Register 注册新用户
// @Summary 注册
// @Description 注册新用户
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseHTTP{}
// @Failure 400 {object} models.ResponseHTTP{}
// @Param user body models.User true "用户名"
// @Router /user/register [post]
func Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Code:    0,
		})
	}

	database.DBConn.Create(&user)

	return ctx.Status(200).JSON(models.ResponseHTTP{
		Success: true,
		Message: "注册成功",
		Data:    nil,
		Code:    1,
	})
}

func GetUser() {

}

func CreateUser(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"status": "success", "message": "Created user", "data": nil})
}

func UpdateUser() {

}

func DeleteUser() {

}
