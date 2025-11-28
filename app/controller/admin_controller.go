package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/request"
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
	_, err = controller.adminService.Create(&admin)
	if err != nil {
		controller.Error(c, 400, fmt.Errorf("注册失败：%w", err).Error())
		return
	}
	dataID := response.DataId{ID: admin.ID}
	controller.Success(c, dataID)
}

func (controller *AdminController) Me(c *gin.Context) {
	adminId, _ := c.Get("admin_id")
	admin, _ := controller.adminService.GetById(adminId.(uint))
	var response response.MeResponse
	copier.Copy(&response, &admin)

	// 补充当前账号对应角色的菜单
	
	controller.Success(c, response)
}

func (controller *AdminController) Page(c *gin.Context) {
	var pagination model.Pagination
	var filter model.AdminFilter

	if err := request.BindAndSetDefaults(c, &pagination); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	_ = c.ShouldBindQuery(&filter)
	total, admins, err := controller.adminService.GetPageByFilter(filter, pagination)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	pageResponse := response.BuildPageResponse[model.Admin, *response.AdminResponse](admins, total, pagination.Page, pagination.PerPage, response.NewAdminResponse)
	controller.Success(c, pageResponse)
}

func (controller *AdminController) Store(c *gin.Context) {
	// 参数处理
	var adminStoreRequest request.AdminStoreRequest
	if err := request.BindAndSetDefaults(c, &adminStoreRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var admin model.Admin
	err := copier.Copy(&admin, &adminStoreRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}
	var roles []model.Role
	err = copier.Copy(&roles, &adminStoreRequest.Roles)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}

	// 密码处理
	if admin.Password != "" {
		// 密码 hash 处理
		admin.Password, err = model.HashPassword(admin.Password)
		if err != nil {
			controller.Error(c, 400, "密码处理失败")
			return
		}
	}

	// service 处理
	_, err = controller.adminService.CreateWithRoles(&admin, roles)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("新增失败：%w", err).Error())
		return
	}
	dataID := response.DataId{ID: admin.ID}
	controller.Success(c, dataID)

}

func (controller *AdminController) Update(c *gin.Context) {
	// 参数处理
	var adminUpdateRequest request.AdminUpdateRequest
	if err := request.BindAndSetDefaults(c, &adminUpdateRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 角色 id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	if id == 1 {
		controller.Error(c, http.StatusBadRequest, "参数异常")
	}

	var admin model.Admin
	err = copier.Copy(&admin, &adminUpdateRequest)
	if err != nil {
		controller.Error(c, 400, "数据获取失败")
		return
	}
	admin.ID = uint(id)

	// 密码处理
	if admin.Password != "" {
		// 密码 hash 处理
		admin.Password, err = model.HashPassword(admin.Password)
		if err != nil {
			controller.Error(c, 400, "密码处理失败")
			return
		}
	}

	var roles []model.Role
	err = copier.Copy(&roles, &adminUpdateRequest.Roles)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}

	// service 处理
	_, err = controller.adminService.UpdateWithRoles(&admin, roles)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("新增失败：%w", err).Error())
		return
	}
	controller.Success(c, nil)
}

func (controller *AdminController) Destroy(c *gin.Context) {
	// admin id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}
	var admin model.Admin
	admin.ID = uint(id)
	err = controller.adminService.Delete(&admin)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}
	controller.Success(c, nil)
}

func (controller *AdminController) Show(c *gin.Context) {
	// admin id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
	}

	admin, err := controller.adminService.GetById(uint(id))
	if err != nil {
		controller.Error(c, http.StatusNotFound, err.Error())
	}

	var adminResponse response.AdminResponse
	copier.Copy(&adminResponse, &admin)

	controller.Success(c, adminResponse)
}
