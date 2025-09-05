package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}

	name := c.Query("name")
	if name != "" {
		filter.Name = &name
	}

	total, roles, err := controller.roleService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse response.PageResponse[response.RoleResponse]
	pageResponse.Total = total
	pageResponse.Page = pagination.Page
	pageResponse.PerPage = pagination.PerPage
	var roleResponses []response.RoleResponse
	err = copier.Copy(&roleResponses, &roles)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	pageResponse.Data = roleResponses
	controller.Success(c, pageResponse)

}

func (controller *RoleController) Store(c *gin.Context) {
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
