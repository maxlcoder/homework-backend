package oms_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/admin/controller"
	"github.com/maxlcoder/homework-backend/app/modules/oms/request"
	"github.com/maxlcoder/homework-backend/app/modules/oms/service"
	base_request "github.com/maxlcoder/homework-backend/app/request"

	"github.com/maxlcoder/homework-backend/app/response"
	"github.com/maxlcoder/homework-backend/model"
)

type OrderController struct {
	controller.BaseController
	// 集成服务
	pickingCarService service.PickingCarService
}

func NewOrderController(pickingCarService service.PickingCarService) *OrderController {
	return &OrderController{
		pickingCarService: pickingCarService,
	}
}

func (controller *OrderController) Store(c *gin.Context) {
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
	dataID := response.DataId{ID: pickingCar.ID}
	controller.Success(c, dataID)

}
