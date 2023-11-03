package user

import (
	"admin-service-go/global"
	"admin-service-go/internal/models"
	"admin-service-go/pkg/app"
	"admin-service-go/pkg/errcode"
	"admin-service-go/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

// Register 注册新用户
// @Summary 注册
// @Description 注册新用户
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserSwagger true "接口入参"
// @Success 200 {object} models.ResponseHTTP{}
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user/register [post]
func Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	response := app.NewResponse(ctx)

	// parse
	err := response.BodyParserErrorResponse(user)
	if err != nil {
		return err
	}

	// validation
	validateErrRes, validateErr := validation.ValidateStruct(user)
	if validateErr != nil {
		return response.ToErrorResponse(errcode.InvalidParams.WithDetails(validateErrRes))
	}

	// 正常的业务逻辑
	global.DBEngine.Create(&user)

	return response.ToResponse(errcode.Success.WithDetails(user))

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
