package validation

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate
var translator ut.Translator

func init() {
	InitValidator()
}

func InitValidator() {
	validate = validator.New()
	uni := ut.New(zh.New())
	// 指定为中文
	translator, _ = uni.GetTranslator("zh")

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("json")
		if name == "-" {
			return ""
		}
		return name
	})

	err := zhtranslations.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		fmt.Println(err)
	}
}

func ValidateStruct(data interface{}) (fiber.Map, error) {
	err := validate.Struct(data)

	errResponse := fiber.Map{}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			// { 'name': '不能为空' }
			errResponse[err.Field()] = err.Translate(translator)
		}
		return errResponse, err
	}
	return errResponse, nil

}
