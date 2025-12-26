package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/modules/core/admin/controller"
	"github.com/maxlcoder/homework-backend/app/modules/wms/admin/request"
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	base_request "github.com/maxlcoder/homework-backend/app/request"

	"github.com/maxlcoder/homework-backend/app/modules/wms/admin/response"
	"github.com/maxlcoder/homework-backend/app/modules/wms/service"
)

type BinController struct {
	controller.BaseController
	// 集成服务
	binService service.BinServiceInterface
}

// GetParamUint 获取uint类型参数
func (controller *BinController) GetParamUint(c *gin.Context, param string) (uint, error) {
	str := c.Param(param)
	id, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func NewBinController(binService service.BinServiceInterface) *BinController {
	return &BinController{
		binService: binService,
	}
}

func (controller *BinController) Page(c *gin.Context) {
	//page := c.DefaultQuery("page", "1")
	//perPage := c.DefaultQuery("per_page", "10")
	//
	//controller.pickingCarService

	controller.Success(c, nil)

}

func (controller *BinController) Show(c *gin.Context) {

	controller.Success(c, nil)

}

func (controller *BinController) Store(c *gin.Context) {
	// 参数处理
	var binStoreRequest request.BinStoreRequest
	if err := base_request.BindAndSetDefaults(c, &binStoreRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var bin model.Bin
	err := copier.Copy(&bin, &binStoreRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}

	// service 处理
	createdBin, err := controller.binService.Create(&bin)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("新增失败：%w", err).Error())
		return
	}

	// 构建响应
	var binResponse response.BinResponse
	err = copier.Copy(&binResponse, createdBin)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "响应构建失败")
		return
	}

	controller.Success(c, binResponse)
}

func (controller *BinController) Update(c *gin.Context) {
	// 参数处理
	var binUpdateRequest request.BinUpdateRequest
	if err := base_request.BindAndSetDefaults(c, &binUpdateRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 获取库位ID
	id, err := controller.GetParamUint(c, "id")
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "无效的库位ID")
		return
	}

	// 查询库位
	bin, err := controller.binService.GetByID(id)
	if err != nil {
		controller.Error(c, http.StatusNotFound, "库位不存在")
		return
	}

	// 更新字段
	err = copier.Copy(bin, &binUpdateRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据更新失败")
		return
	}

	// service 处理
	updatedBin, err := controller.binService.Update(bin)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("更新失败：%w", err).Error())
		return
	}

	// 构建响应
	var binResponse response.BinResponse
	err = copier.Copy(&binResponse, updatedBin)
	if err != nil {
		controller.Error(c, http.StatusInternalServerError, "响应构建失败")
		return
	}

	controller.Success(c, binResponse)
}

func (controller *BinController) Destroy(c *gin.Context) {

	controller.Success(c, nil)

}
