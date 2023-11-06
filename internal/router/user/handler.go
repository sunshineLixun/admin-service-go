package user

import (
	"admin-service-go/global"
	"admin-service-go/internal/models"
	"admin-service-go/pkg/app"
	"admin-service-go/pkg/code"
	"admin-service-go/pkg/validation"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Register 注册新用户
// @Summary 注册
// @Description 注册新用户
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserSwagger true "接口入参"
// @Success 200 {object} models.ResponseHTTP{}
// @Failure 400 {object} models.ResponseHTTP{} "请求错误"
// @Failure 500 {object} models.ResponseHTTP{} "内部错误"
// @Router /api/v1/user/register [post]
func Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	response := app.NewResponse(ctx)

	// parse
	err := response.BodyParserErrorResponse(&user)
	if err != nil {
		return response.InternalServerErrorToResponse(err.Error())
	}

	// 判断是否已经注册过
	isExist := getUserByUserName(user)

	if isExist {
		return response.ToResponse("已经存在该用户", nil)
	}

	// 密码加密
	hash, err := hashPassword(user.Password)
	if err != nil {
		return response.InternalServerErrorToResponse(err.Error())
	}

	user.Password = hash

	// validation
	validateErrRes, validateErr := validation.ValidateStruct(user)
	if validateErr != nil {
		return response.BadRequestToResponse(validateErrRes)
	}

	// 正常的业务逻辑
	if res := global.DBEngine.Create(&user); res.Error != nil {
		return response.InternalServerErrorToResponse(res.Error.Error())
	}

	newUser := models.ResponseUser{
		Model:    user.Model,
		UserName: user.UserName,
	}

	return response.ToResponse(code.Success, newUser)

}

func getUserByUserName(user *models.User) bool {

	res := global.DBEngine.Where("user_name = ?", user.UserName).First(&user)

	return res.RowsAffected > 0

}

// GetAllUser 获取所有用户
// @Summary 获取所有用户
// @Description 获取所有用户
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseHTTP{data=[]models.ResponseUser}
// @Failure 400 {object} models.ResponseHTTP{} "请求错误"
// @Failure 500 {object} models.ResponseHTTP{} "内部错误"
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
// @Failure 400 {object} models.ResponseHTTP{} "请求错误"
// @Failure 500 {object} models.ResponseHTTP{} "内部错误"
// @Router /api/v1/user/{id} [get]
func GetUserById(ctx *fiber.Ctx) error {

	response := app.NewResponse(ctx)

	id := ctx.Params("id")

	user := new(models.User)

	if err := global.DBEngine.First(&user, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ToErrorResponse(http.StatusOK, fmt.Sprintf("未找到id为%s的用户", id), nil)
		}

		return response.InternalServerErrorToResponse(err.Error())

	}

	return response.ToResponse(code.Success, user)
}

func UpdateUser() {

}

// DeleteUser 根据id删除用户
// @Summary 根据id删除用户
// @Description 根据id删除用户
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "用户id"
// @Success 200 {object} models.ResponseHTTP{}
// @Failure 400 {object} models.ResponseHTTP{} "请求错误"
// @Failure 500 {object} models.ResponseHTTP{} "内部错误"
// @Router /api/v1/user/{id} [delete]
func DeleteUser(ctx *fiber.Ctx) error {
	response := app.NewResponse(ctx)

	id := ctx.Params("id")

	user := new(models.User)

	if err := global.DBEngine.First(&user, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ToErrorResponse(http.StatusOK, fmt.Sprintf("未找到id为%s的用户", id), nil)
		}

		return response.InternalServerErrorToResponse(err.Error())

	}

	global.DBEngine.Delete(&user)

	return response.ToResponse(code.Success, nil)

}
