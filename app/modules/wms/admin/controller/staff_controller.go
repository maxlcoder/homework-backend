package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/modules/core/admin/controller"
	"github.com/maxlcoder/homework-backend/app/modules/wms/admin/request"
	"github.com/maxlcoder/homework-backend/app/modules/wms/admin/response"
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	"github.com/maxlcoder/homework-backend/app/modules/wms/service"
	base_request "github.com/maxlcoder/homework-backend/app/request"
	base_response "github.com/maxlcoder/homework-backend/app/response"
)

type StaffController struct {
	controller.BaseController
	staffService service.StaffServiceInterface
}

func NewStaffController(staffService service.StaffServiceInterface) *StaffController {
	return &StaffController{
		staffService: staffService,
	}
}

func (controller *StaffController) Page(c *gin.Context) {
	// 获取分页参数
	var pageRequest base_request.PageRequest

	if err := base_request.BindAndSetDefaults(c, &pageRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 获取员工列表
	staffs, total, err := controller.staffService.Page(pageRequest)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "获取员工列表失败")
		return
	}

	pageResponse := base_response.BuildPageResponse[model.Staff, *response.StaffResponse](staffs, total, pageRequest.Page, pageRequest.PerPage, response.NewStaffResponse)

	controller.Success(c, pageResponse)
}

func (controller *StaffController) Show(c *gin.Context) {
	// 获取分页参数
	var pageRequest base_request.PageRequest

	if err := base_request.BindAndSetDefaults(c, &pageRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 获取员工列表
	staffs, total, err := controller.staffService.Page(pageRequest)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "获取员工列表失败")
		return
	}

	pageResponse := base_response.BuildPageResponse[model.Staff, *response.StaffResponse](staffs, total, pageRequest.Page, pageRequest.PerPage, response.NewStaffResponse)

	controller.Success(c, pageResponse)
}

// Store 创建员工
func (controller *StaffController) Store(c *gin.Context) {
	// 绑定请求参数
	var staffStoreRequest request.StaffStoreRequest
	if err := base_request.BindAndSetDefaults(c, &staffStoreRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 组装 model
	var staff model.Staff
	err := copier.Copy(&staff, &staffStoreRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "创建员工失败")
		return
	}

	// 创建员工
	_, err = controller.staffService.Create(&staff)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "创建员工失败")
		return
	}

	controller.Success(c, staff)
}

// Update 更新员工信息
func (controller *StaffController) Update(c *gin.Context) {
	// 获取员工ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的员工ID")
		return
	}

	// 绑定请求参数
	var staffUpdateRequest request.StaffUpdateRequest
	if err := base_request.BindAndSetDefaults(c, &staffUpdateRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, "参数验证失败")
		return
	}

	// 更新员工信息
	updatedStaff, err := controller.staffService.UpdateStaff(uint(id), staffUpdateRequest)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "更新员工失败")
		return
	}

	controller.Success(c, updatedStaff)
}

// Destroy 删除员工
func (controller *StaffController) Destroy(c *gin.Context) {
	// 获取员工ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的员工ID")
		return
	}

	// 删除员工
	err = controller.staffService.Delete(uint(id))
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "删除员工失败")
		return
	}

	controller.Success(c, nil)
}

// UpdateState 更新员工状态
func (controller *StaffController) UpdateState(c *gin.Context) {
	// 获取员工ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的员工ID")
		return
	}

	// 绑定请求参数
	var stateUpdateRequest struct {
		State int8 `json:"state" binding:"required,gte=0,lte=3"`
	}
	if err := c.ShouldBindJSON(&stateUpdateRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, "参数验证失败")
		return
	}

	// 更新员工状态
	updatedStaff, err := controller.staffService.UpdateStaffState(uint(id), model.StaffState(stateUpdateRequest.State))
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "更新员工状态失败")
		return
	}

	controller.Success(c, updatedStaff)
}
