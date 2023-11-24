package user

import (
	"admin-service-go/global"
	"admin-service-go/internal/models"
	roles2 "admin-service-go/internal/router/roles"
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

func isRegistered(user *models.User) bool {

	res := global.DBEngine.Where("user_name = ?", user.UserName).First(&user)

	return res.RowsAffected > 0

}

func getUserByUserId(id string) (*models.User, error) {
	var user models.User
	// 关联查询
	if err := global.DBEngine.Debug().Preload("Roles").First(&user, id).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "未找到该用户")
	}
	return &user, nil
}

// 设置用户角色关联关系
func associationUserRole(user *models.User) {
	// 查找角色
	if user.RoleIds != nil {
		roles := roles2.GetRoleByIds(user.RoleIds)
		// 关联用户角色
		user.Roles = roles
	}
}

func returnResponseUser(user *models.User) models.ResponseUser {
	var roles []models.ResponseRole

	for _, role := range user.Roles {
		roles = append(roles, models.ResponseRole{
			ID:       role.ID,
			RoleName: role.RoleName,
		})
	}

	newUser := models.ResponseUser{
		ID:       user.ID,
		UserName: user.UserName,
		Roles:    roles,
	}

	return newUser
}

// Create 创建新用户
//
//	@Summary		创建新用户
//	@Description	创建新用户
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.CreateUserInput	true	"接口入参"
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
	isExist := isRegistered(user)

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

	associationUserRole(user)

	// 正常的业务逻辑
	if res := global.DBEngine.Create(user); res.Error != nil {
		return response.InternalServerErrorToResponse(res.Error.Error())
	}

	return response.ToResponse(code.Success, returnResponseUser(user))

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

	// 双层循环，数据量大了之后 这里处理不算太好
	// 如何只返回需要的字段？
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

	if err := global.DBEngine.Debug().Preload("Roles").First(&user, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ToErrorResponse(fiber.StatusOK, fmt.Sprintf("未找到id为%s的用户", id), nil)
		}

		return response.InternalServerErrorToResponse(err.Error())

	}

	var roles []models.ResponseRole

	for _, role := range user.Roles {
		roles = append(roles, models.ResponseRole{
			ID:       role.ID,
			RoleName: role.RoleName,
		})
	}

	newUser := models.ResponseUser{
		ID:       user.ID,
		UserName: user.UserName,
		Roles:    roles,
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
		return response.BadRequestToResponse(err.Error())
	}

	// validation
	validateErrRes, validateErr := validation.ValidateStruct(updateUserInput)

	if validateErr != nil {
		return response.BadRequestToResponse(validateErrRes)
	}

	id := ctx.Params("id")

	user, err := getUserByUserId(id)
	if err != nil {
		return response.ToErrorResponse(fiber.StatusBadRequest, err.Error(), nil)
	}

	if user.UserName == "admin" {
		return response.ToErrorResponse(fiber.StatusBadRequest, "不能修改超管用户", nil)
	}

	user.UserName = updateUserInput.UserName

	// 更新角色关联关系
	if updateUserInput.RoleIds != nil && len(updateUserInput.RoleIds) > 0 {
		// 更新之前解除关联关系
		err := global.DBEngine.Debug().Model(&user).Association("Roles").Unscoped().Clear()
		if err != nil {
			return err
		}
		user.RoleIds = updateUserInput.RoleIds
		associationUserRole(user)
	}

	global.DBEngine.Debug().Updates(&user)

	// 再查询一遍用户信息
	global.DBEngine.Debug().First(&user, id)

	return response.ToResponse(code.Success, returnResponseUser(user))
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

	user, err := getUserByUserId(id)

	if err != nil {
		return response.ToErrorResponse(fiber.StatusBadRequest, err.Error(), nil)
	}

	if user.UserName == "admin" {
		return response.ToErrorResponse(fiber.StatusBadRequest, "不能删除超管用户", nil)
	}

	// 解除关联关系
	global.DBEngine.Debug().Select("Roles").Delete(&user, id)

	return response.ToResponse(code.Success, nil)

}
