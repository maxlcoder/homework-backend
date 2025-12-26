package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/modules/core/admin/response"
	"github.com/maxlcoder/homework-backend/app/modules/core/service"
	base_response "github.com/maxlcoder/homework-backend/app/response"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/validator"
)

type MenuController struct {
	BaseController
	// 集成服务
	menuService service.MenuServiceInterface
}

func NewMenuController(menuService service.MenuServiceInterface) *MenuController {
	return &MenuController{
		menuService: menuService,
	}
}

func (controller *MenuController) Page(c *gin.Context) {
	var pagination model.Pagination
	var filter model.MenuFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)

	total, roles, err := controller.menuService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse base_response.PageResponse[response.RoleResponse]
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

func (controller *MenuController) Store(c *gin.Context) {
	var pagination model.Pagination
	var filter model.MenuFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.menuService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse base_response.PageResponse[response.UserResponse]
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
func (controller *MenuController) Update(c *gin.Context) {
	var pagination model.Pagination
	var filter model.MenuFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.menuService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse base_response.PageResponse[response.UserResponse]
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

func (controller *MenuController) Destroy(c *gin.Context) {
	var pagination model.Pagination
	var filter model.MenuFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.menuService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse base_response.PageResponse[response.UserResponse]
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

func (controller *MenuController) Show(c *gin.Context) {
	var pagination model.Pagination
	var filter model.MenuFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&filter)
	total, users, err := controller.menuService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	var pageResponse base_response.PageResponse[response.UserResponse]
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
