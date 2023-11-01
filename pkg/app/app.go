package app

import (
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

func (r *Response) ToResponse(data interface{}) error {
	if data == nil {
		// 如果没传，给个默认的json
		data = fiber.Map{}
	}
	if err := r.Ctx.Status(http.StatusOK).JSON(data); err != nil {
		newErr := errcode.NewError(100, err.Error())
		return r.ToErrorResponse(newErr)
	}

	return nil
}

func (r *Response) ToErrorResponse(err *errcode.Error) error {
	response := fiber.Map{"code": err.Code(), "msg": err.Msg()}
	data := err.Data()
	if len(data) > 0 {
		response["data"] = data
	}

	return r.Ctx.Status(err.StatusCode()).JSON(response)

}
