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

type PickingBasketController struct {
	controller.BaseController
	// 集成服务
	pickingBasketService service.PickingBasketServiceInterface
}

func NewPickingBasketController(pickingBasketService service.PickingBasketServiceInterface) *PickingBasketController {
	return &PickingBasketController{
		pickingBasketService: pickingBasketService,
	}
}

func (controller *PickingBasketController) Page(c *gin.Context) {
	var pageRequest request.PickingBasketPageRequest
	if err := base_request.BindAndSetDefaults(c, &pageRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 分页查询
	pickingBaskets, count, err := controller.pickingBasketService.Page(pageRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("获取拣货框列表失败：%w", err).Error())
		return
	}
	// 分页响应
	pageResponse := base_response.BuildPageResponse[model.PickingBasket, response.PickingBasketResponse](pickingBaskets, count, pageRequest.Page, pageRequest.PerPage)

	controller.Success(c, pageResponse)

}

func (controller *PickingBasketController) Show(c *gin.Context) {

	// 拣货框 id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "参数转换失败")
	}

	// service 处理
	pickingBasket, err := controller.pickingBasketService.FindById(uint(id))
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("获取拣货框失败：%w", err).Error())
		return
	}

	controller.Success(c, response.ToPickingBasketResponse(*pickingBasket))

}

func (controller *PickingBasketController) Store(c *gin.Context) {
	// 参数处理
	var pickingBasketStoreRequest request.PickingBasketStoreRequest
	if err := base_request.BindAndSetDefaults(c, &pickingBasketStoreRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var pickingBasket model.PickingBasket
	err := copier.Copy(&pickingBasket, &pickingBasketStoreRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}

	// service 处理
	_, err = controller.pickingBasketService.Create(&pickingBasket)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("新增失败：%w", err).Error())
		return
	}
	dataID := base_response.DataId{ID: pickingBasket.ID}
	controller.Success(c, dataID)

}

func (controller *PickingBasketController) Update(c *gin.Context) {
	// 参数处理
	var pickingBasketUpdateRequest request.PickingBasketUpdateRequest
	if err := base_request.BindAndSetDefaults(c, &pickingBasketUpdateRequest); err != nil {
		controller.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var pickingBasket model.PickingBasket
	err := copier.Copy(&pickingBasket, &pickingBasketUpdateRequest)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "数据获取失败")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "参数转换失败")
	}

	pickingBasket.ID = uint(id)

	// service 处理
	_, err = controller.pickingBasketService.Update(&pickingBasket)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("更新失败：%w", err).Error())
		return
	}
	dataID := base_response.DataId{ID: pickingBasket.ID}
	controller.Success(c, dataID)
}

func (controller *PickingBasketController) Destroy(c *gin.Context) {
	// 拣货框 id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		controller.Error(c, http.StatusBadRequest, "参数转换失败")
	}

	// service 处理
	err = controller.pickingBasketService.Delete(uint(id))
	if err != nil {
		controller.Error(c, http.StatusBadRequest, fmt.Errorf("删除失败：%w", err).Error())
		return
	}

	controller.Success(c, nil)
}
