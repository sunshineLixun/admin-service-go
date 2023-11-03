package app

import (
	"admin-service-go/internal/models"
	"admin-service-go/pkg/errcode"
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

func (r *Response) ToResponse(err *errcode.Error) error {
	response := models.ResponseHTTP{
		Success: true,
		Message: err.Msg(),
		Data:    err.Data(),
		Code:    err.Code(),
	}

	return r.Ctx.Status(http.StatusOK).JSON(response)
}

func (r *Response) ToErrorResponse(err *errcode.Error) error {
	response := models.ResponseHTTP{
		Success: false,
		Data:    err.Data(),
		Message: err.Msg(),
		Code:    err.Code(),
	}

	return r.Ctx.Status(http.StatusBadRequest).JSON(response)
}

func (r *Response) BodyParserErrorResponse(out interface{}) error {

	if err := r.Ctx.BodyParser(out); err != nil {
		return r.ToErrorResponse(errcode.NewError(1, err.Error()))
	}

	return nil

}
