package middleware

import (
	"admin-service-go/global"
	"admin-service-go/pkg/app"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(global.JWTSetting.Secret)},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	response := app.NewResponse(c)
	if err.Error() == "Missing or malformed JWT" {
		return response.ToErrorResponse(fiber.StatusBadRequest, "JWT 缺失或格式错误", nil)
	}
	return response.ToErrorResponse(fiber.StatusUnauthorized, "JWT 无效或过期", nil)

}
