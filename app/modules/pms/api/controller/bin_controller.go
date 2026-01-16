package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/modules/core/admin/controller"
	"github.com/maxlcoder/homework-backend/app/modules/wms/service"
)

type BinController struct {
	controller.BaseController
	// 集成服务
	binService service.BinServiceInterface
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
