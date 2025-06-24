package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/response"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/validator"
	"github.com/maxlcoder/homework-backend/service"
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
	var paginationQuery model.PaginationQuery
	var userFilter model.UserFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &paginationQuery); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&userFilter)
	total, users, err := controller.userService.GetPageByFilter(userFilter, paginationQuery)
	log.Println(total, users)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
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
