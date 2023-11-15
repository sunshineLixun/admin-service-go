package jwt

import (
	"admin-service-go/global"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

var (
	UserName = "username"
	UserId   = "user_id"
	Exp      = "exp"
)

func CreateJwtToken(userName string, userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[UserName] = userName
	claims[UserId] = userId
	// 3天后过期
	claims[Exp] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(global.JWTSetting.Secret))

	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidToken(ctx *fiber.Ctx, id string) bool {

	token := ctx.Locals(UserName).(*jwt.Token)

	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := token.Claims.(jwt.MapClaims)

	uid := int(claims[UserId].(float64))

	return uid == n
}
