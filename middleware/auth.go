package middleware

import (
	"admin-service-go/global"
	"admin-service-go/pkg/app"
	"errors"
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
	if errors.Is(err, jwtware.ErrJWTMissingOrMalformed) {
		return response.ToJwtErrorResponse(fiber.StatusBadRequest, "Token失效，请重新登录", nil)
	}

	return response.ToErrorResponse(fiber.StatusUnauthorized, "无权限", nil)

}
