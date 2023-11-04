package user

import (
	"admin-service-go/global"
	"admin-service-go/internal/models"
	"admin-service-go/pkg/app"
	"admin-service-go/pkg/code"
	"admin-service-go/pkg/validation"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
	err := response.BodyParserErrorResponse(&user)
	if err != nil {
		return err
	}

	// validation
	validateErrRes, validateErr := validation.ValidateStruct(user)
	if validateErr != nil {
		return response.BadRequestToResponse(validateErrRes)
	}

	// 正常的业务逻辑
	if res := global.DBEngine.Create(&user); res.Error != nil {
		return response.InternalServerErrorToResponse(res.Error.Error())
	}

	return response.ToResponse(code.Success, user)

}

// GetAllUser 获取所有用户
// @Summary 获取所有用户
// @Description 获取所有用户
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseHTTP{data=[]models.User}
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user [get]
func GetAllUser(ctx *fiber.Ctx) error {

	response := app.NewResponse(ctx)

	var users []models.User

	if res := global.DBEngine.Find(&users); res.Error != nil {
		return response.ToErrorResponse(http.StatusServiceUnavailable, code.ServiceFail, nil)
	}

	return response.ToResponse(code.Success, users)

}

// GetUserById 根据id获取用户详情
// @Summary 根据id获取用户详情
// @Description 根据id获取用户详情
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "用户id"
// @Success 200 {object} models.ResponseHTTP{data=[]models.User}
// @Failure 404 {object} errcode.Error "请求错误"
// @Failure 503 {object} errcode.Error "内部错误"
// @Router /api/v1/user/{id} [get]
func GetUserById(ctx *fiber.Ctx) error {

	response := app.NewResponse(ctx)

	id := ctx.Params("id")

	user := new(models.User)

	if err := global.DBEngine.First(&user, id).Error; err != nil {
		return response.ToErrorResponse(http.StatusOK, fmt.Sprintf("未找到id为%s的用户", id), nil)
	}

	return response.ToResponse(code.Success, user)
}

func UpdateUser() {

}

func DeleteUser() {

}
