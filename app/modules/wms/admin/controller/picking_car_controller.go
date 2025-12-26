package controller

import (
	"fmt"
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

type PickingCarController struct {
	controller.BaseController
	// 集成服务
	pickingCarService service.PickingCarServiceInterface
}

func NewPickingCarController(pickingCarService service.PickingCarServiceInterface) *PickingCarController {
	return &PickingCarController{
		pickingCarService: pickingCarService,
	}
}

func (controller *PickingCarController) Page(c *gin.Context) {
	var pageRequest base_request.PageRequest
	if err := base_request.BindAndSetDefaults(c, &pageRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 分页查询
	pickingCars, count, err := controller.pickingCarService.Page(pageRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("获取拣货车列表失败：%w", err).Error())
		return
	}
	// 分页相应
	pageResponse := base_response.BuildPageResponse[model.PickingCar, *response.PickingCarResponse](pickingCars, count, pageRequest.Page, pageRequest.PerPage, response.NewPickingCarResponse)

	controller.Success(c, pageResponse)

}

func (controller *PickingCarController) Show(c *gin.Context) {

	// 拣货车 id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "参数转换失败")
	}

	// service 处理
	pickingCar, err := controller.pickingCarService.FindById(uint(id))
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("获取拣货车失败：%w", err).Error())
		return
	}

	controller.Success(c, pickingCar)

}

func (controller *PickingCarController) Store(c *gin.Context) {
	// 参数处理
	var pickingCarStoreRequest request.PickingCarStoreRequest
	if err := base_request.BindAndSetDefaults(c, &pickingCarStoreRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var pickingCar model.PickingCar
	err := copier.Copy(&pickingCar, &pickingCarStoreRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}

	// service 处理
	_, err = controller.pickingCarService.Create(&pickingCar)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("新增失败：%w", err).Error())
		return
	}
	dataID := base_response.DataId{ID: pickingCar.ID}
	controller.Success(c, dataID)

}

func (controller *PickingCarController) Update(c *gin.Context) {
	// 参数处理
	var pickingCarUpdateRequest request.PickingCarUpdateRequest
	if err := base_request.BindAndSetDefaults(c, &pickingCarUpdateRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var pickingCar model.PickingCar
	err := copier.Copy(&pickingCar, &pickingCarUpdateRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "参数转换失败")
	}

	pickingCar.ID = uint(id)

	// service 处理
	_, err = controller.pickingCarService.Update(&pickingCar)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("新增失败：%w", err).Error())
		return
	}
	dataID := base_response.DataId{ID: pickingCar.ID}
	controller.Success(c, dataID)
}

func (controller *PickingCarController) Destroy(c *gin.Context) {

	controller.Success(c, nil)

}
