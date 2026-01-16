package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/modules/core/admin/request"
	"github.com/maxlcoder/homework-backend/app/modules/core/admin/response"
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/app/modules/core/service"
	base_request "github.com/maxlcoder/homework-backend/app/request"
	base_response "github.com/maxlcoder/homework-backend/app/response"
	base_model "github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/validator"
)

type AdminController struct {
	BaseController
	// 集成服务
	adminService service.AdminServiceInterface
}

// GetParamUint 获取uint类型参数
func (controller *AdminController) GetParamUint(c *gin.Context, param string) (uint, error) {
	str := c.Param(param)
	id, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
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
	admin.Password, err = base_model.HashPassword(admin.Password)
	if err != nil {
		controller.Error(c, 400, "密码处理失败")
		return
	}
	_, err = controller.adminService.Create(&admin, nil)
	if err != nil {
		controller.Error(c, 400, fmt.Errorf("注册失败：%w", err).Error())
		return
	}
	dataID := base_response.DataId{ID: admin.ID}
	controller.Success(c, dataID)
}

func (controller *AdminController) Me(c *gin.Context) {
	adminId, _ := c.Get("login_admin_id")
	admin, _ := controller.adminService.FindById(adminId.(uint))
	var meResponse response.MeResponse
	copier.Copy(&meResponse, &admin)

	// 补充当前账号对应角色的菜单
	// 获取角色全部菜单
	menus, _ := controller.adminService.GetMenusWithChildrenByRoleId(admin.RoleId)
	// 菜单 -> tree
	meResponse.Menus = response.TreesToResponse(menus)
	controller.Success(c, meResponse)
}

func (controller *AdminController) Page(c *gin.Context) {
	var pageRequest request.AdminPageRequest
	if err := base_request.BindAndSetDefaults(c, &pageRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 分页查询
	admins, count, err := controller.adminService.Page(pageRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("获取管理员列表失败：%w", err).Error())
		return
	}

	// 分页响应
	pageResponse := base_response.BuildPageResponse[model.Admin, response.AdminResponse](admins, count, pageRequest.Page, pageRequest.PerPage)

	controller.Success(c, pageResponse)
}

func (controller *AdminController) Store(c *gin.Context) {
	// 参数处理
	var adminStoreRequest request.AdminStoreRequest
	if err := base_request.BindAndSetDefaults(c, &adminStoreRequest); err != nil {
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
		admin.Password, err = base_model.HashPassword(admin.Password)
		if err != nil {
			controller.Error(c, 400, "密码处理失败")
			return
		}
	}

	// service 处理
	createdAdmin, err := controller.adminService.Create(&admin, roles)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("新增失败：%w", err).Error())
		return
	}

	// 构建响应
	var adminResponse response.AdminResponse
	err = copier.Copy(&adminResponse, createdAdmin)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "响应构建失败")
		return
	}

	controller.Success(c, adminResponse)
}

func (controller *AdminController) Update(c *gin.Context) {
	// 参数处理
	var adminUpdateRequest request.AdminUpdateRequest
	if err := base_request.BindAndSetDefaults(c, &adminUpdateRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 获取管理员ID
	id, err := controller.GetParamUint(c, "id")
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的管理员ID")
		return
	}

	if id == 1 {
		controller.Error(c, http.StatusBadRequest, "参数异常")
		return
	}

	// 查询管理员
	admin, err := controller.adminService.FindById(id)
	if err != nil {
		controller.Error(c, http.StatusNotFound, "管理员不存在")
		return
	}

	// 更新字段
	err = copier.Copy(admin, &adminUpdateRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据更新失败")
		return
	}

	var roles []model.Role
	err = copier.Copy(&roles, &adminUpdateRequest.Roles)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}

	// 密码处理
	if admin.Password != "" {
		// 密码 hash 处理
		admin.Password, err = base_model.HashPassword(admin.Password)
		if err != nil {
			controller.Error(c, 400, "密码处理失败")
			return
		}
	}

	// service 处理
	updatedAdmin, err := controller.adminService.Update(admin, roles)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("更新失败：%w", err).Error())
		return
	}

	// 构建响应
	var adminResponse response.AdminResponse
	err = copier.Copy(&adminResponse, updatedAdmin)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "响应构建失败")
		return
	}

	controller.Success(c, adminResponse)
}

func (controller *AdminController) Destroy(c *gin.Context) {
	// 获取管理员ID
	id, err := controller.GetParamUint(c, "id")
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的管理员ID")
		return
	}

	if id == 1 {
		controller.Error(c, http.StatusBadRequest, "参数异常")
		return
	}

	err = controller.adminService.Delete(id)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("删除失败：%w", err).Error())
		return
	}

	controller.Success(c, nil)
}

func (controller *AdminController) Show(c *gin.Context) {
	// 获取管理员ID
	id, err := controller.GetParamUint(c, "id")
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的管理员ID")
		return
	}

	admin, err := controller.adminService.FindById(id)
	if err != nil {
		controller.Error(c, http.StatusNotFound, err.Error())
		return
	}

	var adminResponse response.AdminResponse
	copier.Copy(&adminResponse, &admin)

	controller.Success(c, adminResponse)
}
