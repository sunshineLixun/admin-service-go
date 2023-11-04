package app

import (
	"admin-service-go/internal/models"
	"admin-service-go/pkg/code"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Response struct {
	Ctx *fiber.Ctx
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *fiber.Ctx) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(msg string, data interface{}) error {
	response := models.ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    data,
		Code:    1,
	}

	return r.Ctx.Status(http.StatusOK).JSON(response)
}

func (r *Response) ToErrorResponse(status int, msg string, data interface{}) error {
	response := models.ResponseHTTP{
		Success: false,
		Data:    data,
		Message: msg,
		Code:    0,
	}

	return r.Ctx.Status(status).JSON(response)
}

func (r *Response) BadRequestToResponse(data interface{}) error {
	return r.ToErrorResponse(http.StatusBadRequest, code.ParamsFail, data)
}

func (r *Response) InternalServerErrorToResponse(data interface{}) error {
	return r.ToErrorResponse(http.StatusInternalServerError, code.ServiceFail, data)
}

func (r *Response) BodyParserErrorResponse(out interface{}) error {

	if err := r.Ctx.BodyParser(&out); err != nil {
		return r.InternalServerErrorToResponse(err.Error())
	}

	return nil

}
