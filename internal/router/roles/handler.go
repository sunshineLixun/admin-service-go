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

func getRoleById(roleId string) (*models.Role, bool) {

	var role models.Role

	db := global.DBEngine.Debug().Raw("select * from roles where id = ?", roleId).Scan(&role)

	return &role, db.RowsAffected > 0
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

// GetAllRoles 获取所有角色
//
//	@Summary		获取所有角色
//	@Description	获取所有角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.ResponseHTTP{}
//	@Failure		400		{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500		{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/roles/getAllRoles [get]
func GetAllRoles(c *fiber.Ctx) error {

	response := app.NewResponse(c)

	var roles []models.ResponseRole

	global.DBEngine.Debug().Raw("select id, role_name from roles").Scan(&roles)

	return response.ToResponse(code.Success, roles)
}

// GetRoleByRoleId 根据角色id获取角色详情
//
//	@Summary		根据角色id获取角色详情
//	@Description	根据角色id获取角色详情
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.ResponseHTTP{}
//	@Failure		400		{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500		{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/roles/getAllRoles [get]
func GetRoleByRoleId(c *fiber.Ctx) error {

	response := app.NewResponse(c)

	roleId := c.Params("id")

	var role models.ResponseRole

	_, isExist := getRoleById(roleId)

	if !isExist {
		return response.ToErrorResponse(fiber.StatusBadRequest, "不存在该角色", nil)
	}

	global.DBEngine.Debug().Raw("select id, role_name from roles where id = ?", roleId).Scan(&role)

	return response.ToResponse(code.Success, role)
}

// UpdateRole 修改角色信息
//
//	@Summary		修改角色信息
//	@Description	修改角色信息
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int						true	"角色id"
//	@Param			user	body		models.UpdateRoleInput	true	"接口入参"
//	@Success		200		{object}	models.ResponseHTTP{data=models.User}
//	@Failure		400		{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500		{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/role/{id} [patch]
func UpdateRole(ctx *fiber.Ctx) error {
	response := app.NewResponse(ctx)
	var updateRoleInput models.UpdateRoleInput

	err := response.BodyParserErrorResponse(&updateRoleInput)
	if err != nil {
		return response.InternalServerErrorToResponse(err.Error())
	}

	// validation
	validateErrRes, validateErr := validation.ValidateStruct(updateRoleInput)

	if validateErr != nil {
		return response.BadRequestToResponse(validateErrRes)
	}

	roleId := ctx.Params("id")

	_, isExist := getRoleById(roleId)

	if !isExist {
		return response.ToErrorResponse(fiber.StatusBadRequest, "不存在该角色", nil)
	}

	role := new(models.Role)

	global.DBEngine.Debug().Raw("update roles set role_name = ? where id = ?", updateRoleInput.RoleName, roleId)

	role, _ = getRoleById(roleId)

	newRole := models.ResponseRole{
		ID:       role.ID,
		RoleName: role.RoleName,
	}

	return response.ToResponse(code.Success, newRole)
}

// DeleteRole 删除角色信息
//
//	@Summary		删除角色信息
//	@Description	删除角色信息
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int						true	"角色id"
//	@Success		200		{object}	models.ResponseHTTP{data=models.User}
//	@Failure		400		{object}	models.ResponseHTTP{}	"请求错误"
//	@Failure		500		{object}	models.ResponseHTTP{}	"内部错误"
//	@Router			/api/v1/role/{id} [patch]
func DeleteRole(c *fiber.Ctx) error {
	response := app.NewResponse(c)

	roleId := c.Params("id")

	_, isExist := getRoleById(roleId)

	if !isExist {
		return response.ToErrorResponse(fiber.StatusBadRequest, "不存在该角色", nil)
	}

	if err := global.DBEngine.Debug().Delete(&models.Role{}, roleId).Error; err != nil {
		return response.InternalServerErrorToResponse(err.Error())
	}

	return response.ToResponse(code.Success, nil)
}
