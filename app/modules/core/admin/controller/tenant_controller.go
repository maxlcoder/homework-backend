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
)

type TenantController struct {
	BaseController
	// 集成服务
	tenantService service.TenantServiceInterface
}

// GetParamUint 获取uint类型参数
func (controller *TenantController) GetParamUint(c *gin.Context, param string) (uint, error) {
	str := c.Param(param)
	id, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func NewTenantController(tenantService service.TenantServiceInterface) *TenantController {
	return &TenantController{
		tenantService: tenantService,
	}
}

func (controller *TenantController) Page(c *gin.Context) {

	var pageRequest request.TenantPageRequest
	if err := base_request.BindAndSetDefaults(c, &pageRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 分页查询
	tenants, count, err := controller.tenantService.Page(pageRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("获取租户列表失败：%w", err).Error())
		return
	}
	// 分页响应
	pageResponse := base_response.BuildPageResponse[model.Tenant, response.TenantResponse](tenants, count, pageRequest.Page, pageRequest.PerPage)

	controller.Success(c, pageResponse)

}

func (controller *TenantController) Show(c *gin.Context) {
	id, err := controller.GetParamUint(c, "id")
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的租户ID")
		return
	}

	tenant, err := controller.tenantService.FindById(id)
	if err != nil {
		controller.Error(c, http.StatusNotFound, err.Error())
		return
	}

	var tenantResponse response.TenantResponse
	err = copier.Copy(&tenantResponse, tenant)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "响应构建失败")
		return
	}

	controller.Success(c, tenantResponse)
}

func (controller *TenantController) Store(c *gin.Context) {
	// 参数处理
	var tenantStoreRequest request.TenantStoreRequest
	if err := base_request.BindAndSetDefaults(c, &tenantStoreRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var tenant model.Tenant
	err := copier.Copy(&tenant, &tenantStoreRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}

	// service 处理
	createdTenant, err := controller.tenantService.Create(&tenant)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("新增失败：%w", err).Error())
		return
	}

	// 构建响应
	var tenantResponse response.TenantResponse
	err = copier.Copy(&tenantResponse, createdTenant)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "响应构建失败")
		return
	}

	controller.Success(c, tenantResponse)
}

func (controller *TenantController) Update(c *gin.Context) {
	// 参数处理
	var tenantUpdateRequest request.TenantUpdateRequest
	if err := base_request.BindAndSetDefaults(c, &tenantUpdateRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 获取租户ID
	id, err := controller.GetParamUint(c, "id")
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的租户ID")
		return
	}

	// 查询租户
	tenant, err := controller.tenantService.FindById(id)
	if err != nil {
		controller.Error(c, http.StatusNotFound, "租户不存在")
		return
	}

	// 更新字段
	err = copier.Copy(tenant, &tenantUpdateRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据更新失败")
		return
	}

	// service 处理
	updatedTenant, err := controller.tenantService.Update(tenant)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("更新失败：%w", err).Error())
		return
	}

	// 构建响应
	var tenantResponse response.TenantResponse
	err = copier.Copy(&tenantResponse, updatedTenant)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "响应构建失败")
		return
	}

	controller.Success(c, tenantResponse)
}

func (controller *TenantController) Destroy(c *gin.Context) {
	id, err := controller.GetParamUint(c, "id")
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的租户ID")
		return
	}

	err = controller.tenantService.Delete(id)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("删除失败：%w", err).Error())
		return
	}

	controller.Success(c, nil)
}
