package roles

import (
	"admin-service-go/global"
	"admin-service-go/internal/models"
	"admin-service-go/pkg/app"
	"admin-service-go/pkg/code"
	"admin-service-go/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

func getRoleByRoleName(role *models.Role) bool {

	sql := "SELECT * FROM roles WHERE role_name = @role_name"
	// 这么写可以防止sql注入
	res := global.DBEngine.Raw(sql, map[string]interface{}{"role_name": role.RoleName}).Scan(role)
	return res.RowsAffected > 0

}

// CreateRole 创建新角色
//
//	@Summary		创建新角色
//	@Description	创建新角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			role	body		models.InputRole	true	"接口入参"
//	@Success		200		{object}	models.ResponseHTTP{}
//	@Failure		400		{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500		{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/roles/create [post]
func CreateRole(c *fiber.Ctx) error {

	role := new(models.Role)
	response := app.NewResponse(c)

	// parse
	err := response.BodyParserErrorResponse(&role)
	if err != nil {
		return response.InternalServerErrorToResponse(err.Error())
	}

	// validation
	validateErrRes, validateErr := validation.ValidateStruct(role)
	if validateErr != nil {
		return response.BadRequestToResponse(validateErrRes)
	}

	// 判断是否已经创建过
	isExist := getRoleByRoleName(role)
	if isExist {
		return response.ToErrorResponse(fiber.StatusBadRequest, "已经存在该角色", nil)
	}

	// 创建角色
	if res := global.DBEngine.Create(&role); res.Error != nil {
		return response.InternalServerErrorToResponse(res.Error.Error())
	}

	return response.ToResponse(code.Success, nil)
}

func GetAllRoles(c *fiber.Ctx) error {
	return c.SendString("GetAllRoles")
}
