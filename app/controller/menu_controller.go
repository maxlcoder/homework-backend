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

type MenuController struct {
	BaseController
	// 集成服务
	menuService service.RoleServiceInterface
}

func NewMenuController(roleService service.RoleServiceInterface) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (controller *RoleController) Page(c *gin.Context) {
	var paginationQuery model.PaginationQuery
	var filter model.RoleFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &paginationQuery); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, roles, err := controller.roleService.GetPageByFilter(filter, paginationQuery)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse response.PageResponse[response.RoleResponse]
	pageResponse.Total = total
	pageResponse.Page = paginationQuery.Page
	pageResponse.PerPage = paginationQuery.PerPage
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
	var paginationQuery model.PaginationQuery
	var filter model.RoleFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &paginationQuery); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.roleService.GetPageByFilter(filter, paginationQuery)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse response.PageResponse[response.UserResponse]
	pageResponse.Total = total
	pageResponse.Page = paginationQuery.Page
	pageResponse.PerPage = paginationQuery.PerPage
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.ToUserResponse(user))
	}
	pageResponse.Data = userResponses

	controller.Success(c, pageResponse)

}
func (controller *RoleController) Update(c *gin.Context) {
	var paginationQuery model.PaginationQuery
	var filter model.RoleFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &paginationQuery); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.roleService.GetPageByFilter(filter, paginationQuery)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse response.PageResponse[response.UserResponse]
	pageResponse.Total = total
	pageResponse.Page = paginationQuery.Page
	pageResponse.PerPage = paginationQuery.PerPage
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.ToUserResponse(user))
	}
	pageResponse.Data = userResponses

	controller.Success(c, pageResponse)

}

func (controller *RoleController) Destroy(c *gin.Context) {
	var paginationQuery model.PaginationQuery
	var filter model.RoleFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &paginationQuery); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.roleService.GetPageByFilter(filter, paginationQuery)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse response.PageResponse[response.UserResponse]
	pageResponse.Total = total
	pageResponse.Page = paginationQuery.Page
	pageResponse.PerPage = paginationQuery.PerPage
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.ToUserResponse(user))
	}
	pageResponse.Data = userResponses

	controller.Success(c, pageResponse)

}

func (controller *RoleController) Show(c *gin.Context) {
	var paginationQuery model.PaginationQuery
	var filter model.RoleFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &paginationQuery); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.roleService.GetPageByFilter(filter, paginationQuery)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse response.PageResponse[response.UserResponse]
	pageResponse.Total = total
	pageResponse.Page = paginationQuery.Page
	pageResponse.PerPage = paginationQuery.PerPage
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.ToUserResponse(user))
	}
	pageResponse.Data = userResponses

	controller.Success(c, pageResponse)

}
