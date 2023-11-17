package jwt

import (
	"admin-service-go/global"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

var (
	User     = "user"
	UserName = "username"
	UserId   = "user_id"
	Exp      = "exp"
)

func CreateJwtToken(userName string, userId uint) (string, error) {

	claims := jwt.MapClaims{
		UserName: userName,
		UserId:   userId,
		Exp:      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(global.JWTSetting.Secret))

	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidToken(ctx *fiber.Ctx, id string) bool {

	user := ctx.Locals(User).(*jwt.Token)

	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := user.Claims.(jwt.MapClaims)

	uid := int(claims[UserId].(float64))

	return uid == n
}
