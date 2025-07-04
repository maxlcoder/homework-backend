package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	request "github.com/maxlcoder/homework-backend/app/request/admin"
	"github.com/maxlcoder/homework-backend/app/response"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/validator"
	"github.com/maxlcoder/homework-backend/service"
)

type AdminController struct {
	BaseController
	// 集成服务
	adminService service.AdminServiceInterface
}

func NewAdminController(adminService service.AdminServiceInterface) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

func (controller *AdminController) Register(c *gin.Context) {

	var adminCreateRequest request.AdminCreateRequest

	// 进行参数校验
	if errs, ok := validator.BindAndValidateFirst(c, &adminCreateRequest); !ok {
		controller.Error(c, 400, errs)
		return
	}

	var admin model.Admin
	err := copier.Copy(&admin, &adminCreateRequest)
	if err != nil {
		controller.Error(c, 400, "数据获取失败")
		return
	}
	// 密码 hash 处理
	admin.Password, err = model.HashPassword(admin.Password)
	if err != nil {
		controller.Error(c, 400, "密码处理失败")
		return
	}
	fmt.Println(admin.Password)
	_, err = controller.adminService.Create(&admin)
	if err != nil {
		controller.Error(c, 400, fmt.Errorf("注册失败：%w", err).Error())
		return
	}
	dataID := response.DataId{ID: int(admin.ID)}
	controller.Success(c, dataID)
}

func (controller *AdminController) Me(c *gin.Context) {
	controller.adminService.Create(&model.Admin{})
	controller.Success(c, "ttt")
}

func (controller *AdminController) Page(c *gin.Context) {
	var paginationQuery model.PaginationQuery
	var userFilter model.UserFilter
	if err, ok := validator.BindQueryAndValidateAll(c, &paginationQuery); !ok {
		controller.Error(c, http.StatusBadRequest, err)
		return
	}
	_ = c.ShouldBindJSON(&userFilter)
	total, users, err := controller.adminService.GetPageByFilter(userFilter, paginationQuery)
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
