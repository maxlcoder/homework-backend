package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/modules/core/admin/response"
	model2 "github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/app/modules/core/service"
	base_response "github.com/maxlcoder/homework-backend/app/response"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/validator"
)

type AdminUserController struct {
	BaseController
	// 集成服务
	adminService service.AdminServiceInterface
	userService  service.UserServiceInterface
}

func NewAdminUserController(adminService service.AdminServiceInterface, userService service.UserServiceInterface) *AdminUserController {
	return &AdminUserController{
		adminService: adminService,
		userService:  userService,
	}
}

func (controller *AdminUserController) Page(c *gin.Context) {
	var pagination model.Pagination
	var userFilter model2.UserFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &pagination); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&userFilter)
	total, users, err := controller.userService.GetPageByFilter(userFilter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
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
