package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/request"
	"github.com/maxlcoder/homework-backend/app/response"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/validator"
	"github.com/maxlcoder/homework-backend/service"
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

	if err := request.BindAndSetDefaults(c, &pagination); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	_ = c.ShouldBindQuery(&filter)
	total, roles, err := controller.roleService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	pageResponse := response.BuildPageResponse[model.Role, *response.RoleResponse](roles, total, pagination.Page, pagination.PerPage, response.NewRoleResponse)
	controller.Success(c, pageResponse)

}

func (controller *RoleController) Store(c *gin.Context) {
	var roleStoreRequest request.RoleStoreRequest

	if err := request.BindAndSetDefaults(c, &roleStoreRequest); err != nil {
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
	dataID := response.DataId{ID: int(role.ID)}
	controller.Success(c, dataID)
}

func (controller *RoleController) Update(c *gin.Context) {
	var pagination model.Pagination
	var filter model.RoleFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.roleService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse response.PageResponse[response.UserResponse]
	pageResponse.Total = total
	pageResponse.Page = pagination.Page
	pageResponse.PerPage = pagination.PerPage
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.ToUserResponse(user))
	}
	pageResponse.Data = userResponses

	controller.Success(c, pageResponse)

}

func (controller *RoleController) Destroy(c *gin.Context) {
	var pagination model.Pagination
	var filter model.RoleFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.roleService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse response.PageResponse[response.UserResponse]
	pageResponse.Total = total
	pageResponse.Page = pagination.Page
	pageResponse.PerPage = pagination.PerPage
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.ToUserResponse(user))
	}
	pageResponse.Data = userResponses

	controller.Success(c, pageResponse)

}

func (controller *RoleController) Show(c *gin.Context) {
	var pagination model.Pagination
	var filter model.RoleFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.roleService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse response.PageResponse[response.UserResponse]
	pageResponse.Total = total
	pageResponse.Page = pagination.Page
	pageResponse.PerPage = pagination.PerPage
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.ToUserResponse(user))
	}
	pageResponse.Data = userResponses

	controller.Success(c, pageResponse)

}
