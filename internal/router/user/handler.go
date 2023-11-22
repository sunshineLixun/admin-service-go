package user

import (
	"admin-service-go/global"
	"admin-service-go/internal/models"
	"admin-service-go/pkg/app"
	"admin-service-go/pkg/code"
	"admin-service-go/pkg/jwt"
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

func getUserByUserName(user *models.User) bool {

	res := global.DBEngine.Where("user_name = ?", user.UserName).First(&user)

	return res.RowsAffected > 0

}

func getRoleByRoleId(roleId uint) models.Role {

	var role models.Role

	global.DBEngine.First(&role, roleId)

	return role

}

// Create 创建新用户
//
//	@Summary		创建新用户
//	@Description	创建新用户
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserSwagger	true	"接口入参"
//	@Success		200		{object}	models.ResponseHTTP{}
//	@Failure		400		{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500		{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/user/register [post]
func Create(ctx *fiber.Ctx) error {
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
		return response.ToErrorResponse(fiber.StatusBadRequest, "已经存在该用户", nil)
	}

	// validation
	validateErrRes, validateErr := validation.ValidateStruct(user)
	if validateErr != nil {
		return response.BadRequestToResponse(validateErrRes)
	}

	// 密码加密
	hash, err := hashPassword(user.Password)
	if err != nil {
		return response.InternalServerErrorToResponse(err.Error())
	}

	user.Password = hash

	// 查找角色
	if user.RoleIds != nil {
		var roles []models.Role
		for _, v := range user.RoleIds {
			roles = append(roles, getRoleByRoleId(v))
		}
		// 关联用户角色
		user.Roles = roles
	}

	// 正常的业务逻辑
	if res := global.DBEngine.Create(&user); res.Error != nil {
		return response.InternalServerErrorToResponse(res.Error.Error())
	}

	newUser := models.ResponseUser{
		ID:       user.ID,
		UserName: user.UserName,
	}

	return response.ToResponse(code.Success, newUser)

}

// GetAllUser 获取所有用户
//
//	@Summary		获取所有用户
//	@Description	获取所有用户
//	@Tags			用户
//	@Accept			json
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	models.ResponseHTTP{data=[]models.ResponseUser}
//	@Failure		400	{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500	{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/user/getAllUsers [get]
func GetAllUser(ctx *fiber.Ctx) error {

	response := app.NewResponse(ctx)

	var users []models.User
	var responseUsers []models.ResponseUser

	if res := global.DBEngine.Debug().Model(&models.User{}).Preload("Roles").Find(&users); res.Error != nil {
		return response.ToErrorResponse(http.StatusServiceUnavailable, code.ServiceFail, nil)
	}

	for _, user := range users {
		var roles []models.ResponseRole
		for _, role := range user.Roles {
			roles = append(roles, models.ResponseRole{
				ID:       role.ID,
				RoleName: role.RoleName,
			})
		}

		responseUsers = append(responseUsers, models.ResponseUser{
			ID:       user.ID,
			UserName: user.UserName,
			Roles:    roles,
		})
	}

	return response.ToResponse(code.Success, responseUsers)

}

// GetUserById 根据id获取用户详情
//
//	@Summary		根据id获取用户详情
//	@Description	根据id获取用户详情
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"用户id"
//	@Success		200	{object}	models.ResponseHTTP{data=[]models.User}
//	@Failure		400	{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500	{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/user/{id} [get]
func GetUserById(ctx *fiber.Ctx) error {

	response := app.NewResponse(ctx)

	id := ctx.Params("id")

	user := new(models.User)

	if err := global.DBEngine.First(&user, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ToErrorResponse(fiber.StatusOK, fmt.Sprintf("未找到id为%s的用户", id), nil)
		}

		return response.InternalServerErrorToResponse(err.Error())

	}

	newUser := models.ResponseUser{
		ID:       user.ID,
		UserName: user.UserName,
	}

	return response.ToResponse(code.Success, newUser)
}

// UpdateUser 修改用户信息
//
//	@Summary		修改用户信息
//	@Description	修改用户信息
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int						true	"用户id"
//	@Param			user	body		models.UpdateUserInput	true	"接口入参"
//	@Success		200		{object}	models.ResponseHTTP{data=models.User}
//	@Failure		400		{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500		{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/user/{id} [patch]
func UpdateUser(ctx *fiber.Ctx) error {

	response := app.NewResponse(ctx)
	var updateUserInput models.UpdateUserInput

	err := response.BodyParserErrorResponse(&updateUserInput)
	if err != nil {
		return response.InternalServerErrorToResponse(err.Error())
	}

	// validation
	validateErrRes, validateErr := validation.ValidateStruct(updateUserInput)

	if validateErr != nil {
		return response.BadRequestToResponse(validateErrRes)
	}

	id := ctx.Params("id")

	var user models.User

	if !jwt.ValidToken(ctx, id) {
		return response.ToErrorResponse(fiber.StatusInternalServerError, "未找到改用户", nil)
	}

	user.UserName = updateUserInput.UserName

	global.DBEngine.Save(&user)

	newUser := models.ResponseUser{
		ID:       user.ID,
		UserName: user.UserName,
	}

	return response.ToResponse(code.Success, newUser)
}

// DeleteUser 根据id删除用户
//
//	@Summary		根据id删除用户
//	@Description	根据id删除用户
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"用户id"
//	@Success		200	{object}	models.ResponseHTTP{}
//	@Failure		400	{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500	{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/user/{id} [delete]
func DeleteUser(ctx *fiber.Ctx) error {
	response := app.NewResponse(ctx)

	id := ctx.Params("id")

	if !jwt.ValidToken(ctx, id) {
		return response.ToErrorResponse(fiber.StatusInternalServerError, "未找到改用户", nil)
	}

	user := new(models.User)

	if err := global.DBEngine.First(&user, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ToErrorResponse(fiber.StatusOK, fmt.Sprintf("未找到id为%s的用户", id), nil)
		}

		return response.InternalServerErrorToResponse(err.Error())

	}

	global.DBEngine.Delete(&user)

	return response.ToResponse(code.Success, nil)

}
