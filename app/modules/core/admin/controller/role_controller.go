package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/modules/core/admin/request"
	"github.com/maxlcoder/homework-backend/app/modules/core/admin/response"
	"github.com/maxlcoder/homework-backend/app/modules/core/service"
	base_request "github.com/maxlcoder/homework-backend/app/request"
	base_response "github.com/maxlcoder/homework-backend/app/response"
	"github.com/maxlcoder/homework-backend/model"
)

type RoleController struct {
	BaseController
	// 集成服务
	roleService service.RoleServiceInterface
}

func NewRoleController(roleService service.RoleServiceInterface) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (controller *RoleController) Page(c *gin.Context) {
	var pagination model.Pagination
	var filter model.RoleFilter

	if err := base_request.BindAndSetDefaults(c, &pagination); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	_ = c.ShouldBindQuery(&filter)
	total, roles, err := controller.roleService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	pageResponse := base_response.BuildPageResponse[model.Role, response.RoleResponse](roles, total, pagination.Page, pagination.PerPage)
	controller.Success(c, pageResponse)

}

func (controller *RoleController) Store(c *gin.Context) {
	var roleStoreRequest request.RoleStoreRequest

	if err := base_request.BindAndSetDefaults(c, &roleStoreRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var role model.Role
	err := copier.Copy(&role, &roleStoreRequest)
	if err != nil {
		controller.Error(c, 400, "数据获取失败")
		return
	}
	var menus []model.Menu
	err = copier.Copy(&menus, &roleStoreRequest.Menus)
	if err != nil {
		controller.Error(c, 400, "数据获取失败")
		return
	}

	_, err = controller.roleService.CreateWithMenus(&role, menus)
	if err != nil {
		controller.Error(c, 400, fmt.Errorf("新增失败：%w", err).Error())
		return
	}
	dataID := base_response.DataId{ID: role.ID}
	controller.Success(c, dataID)
}

func (controller *RoleController) Update(c *gin.Context) {
	var roleStoreRequest request.RoleStoreRequest

	if err := base_request.BindAndSetDefaults(c, &roleStoreRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 角色 id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var role model.Role
	err = copier.Copy(&role, &roleStoreRequest)
	if err != nil {
		controller.Error(c, 400, "数据获取失败")
		return
	}

	role.ID = uint(id)

	var menus []model.Menu
	err = copier.Copy(&menus, &roleStoreRequest.Menus)
	if err != nil {
		controller.Error(c, 400, "数据获取失败")
		return
	}

	_, err = controller.roleService.UpdateWithMenus(&role, menus)
	if err != nil {
		controller.Error(c, 400, fmt.Errorf("新增失败：%w", err).Error())
		return
	}
	controller.Success(c, nil)

}

func (controller *RoleController) Destroy(c *gin.Context) {
	// 角色 id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var role model.Role
	role.ID = uint(id)
	err = controller.roleService.Delete(&role)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}
	controller.Success(c, nil)

}

func (controller *RoleController) Show(c *gin.Context) {
	// 角色 id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	role, err := controller.roleService.GetById(uint(id))
	if err != nil {
		controller.Error(c, http.StatusNotFound, err.Error())
	}

	var roleResponse response.RoleResponse
	copier.Copy(&roleResponse, &role)

	controller.Success(c, roleResponse)

}
