package validation

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var validate *validator.Validate
var translator ut.Translator

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func init() {
	InitValidator()
}

func InitValidator() {
	validate = validator.New()
	uni := ut.New(zh.New())
	translator, _ = uni.GetTranslator("zh")

	err := zh_translations.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		fmt.Println(err)
	}
}

func ValidateStruct(data interface{}) error {
	err := validate.Struct(data)
	strSli := make([]string, 0)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			strSli = append(strSli, err.Translate(translator))
		}
		return &fiber.Error{
			Code:    0,
			Message: strings.Join(strSli, ","),
		}
	}
	return nil

}
