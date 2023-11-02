package user

import (
	"admin-service-go/global"
	"admin-service-go/internal/models"
	"admin-service-go/pkg/app"
	"admin-service-go/pkg/errcode"
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

	global.DBEngine.Create(&user)

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
	global.Logger.Infof("%s: test/%s", "test", "test-service")
	return app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
}

func UpdateUser() {

}

func DeleteUser() {

}
