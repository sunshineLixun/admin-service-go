package auth

import (
	"admin-service-go/global"
	"admin-service-go/internal/models"
	"admin-service-go/pkg/app"
	jwt2 "admin-service-go/pkg/jwt"
	"admin-service-go/pkg/validation"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByUsername(username string) (*models.User, error) {
	var user models.User

	if err := global.DBEngine.Where("user_name = ?", username).First(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("未找到用户名为%s 的用户", username))
		}
		return nil, err
	}

	return &user, nil
}

// Login 登录
//
//	@Summary		登录
//	@Description	登录
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserSwagger	true	"接口入参"
//	@Success		200		{object}	models.ResponseHTTP{data=string}
//	@Failure		400		{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500		{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {

	response := app.NewResponse(c)

	input := new(models.UserSwagger)

	// 参数解析
	err := response.BodyParserErrorResponse(&input)

	if err != nil {
		return response.ToErrorResponse(fiber.StatusBadRequest, err.Error(), nil)
	}

	// validation
	validateErrRes, validateErr := validation.ValidateStruct(input)
	if validateErr != nil {
		return response.BadRequestToResponse(validateErrRes)
	}

	password := input.Password
	userModel, err := new(models.User), *new(error)

	userModel, err = getUserByUsername(input.UserName)

	if err != nil {
		return response.ToErrorResponse(fiber.StatusUnauthorized, err.Error(), nil)
	}

	if !CheckPasswordHash(password, userModel.Password) {
		return response.ToErrorResponse(fiber.StatusUnauthorized, "密码错误", nil)
	}

	t, err := jwt2.CreateJwtToken(userModel.UserName, userModel.ID)

	if err != nil {
		return response.ToErrorResponse(fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.ToResponse("登录成功", t)

}
