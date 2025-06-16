package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/service"
)

type AdminUserController struct {
	BaseController
	// 注入不用的服务
	userService service.UserServiceInterface
}

func NewAdminUserController(userService service.UserServiceInterface) *AdminUserController {
	return &AdminUserController{
		userService: userService,
	}
}

func (controller *AdminUserController) Register(c *gin.Context) {
	controller.userService.Create(&model.User{})
	controller.Success(c, "ttt")
}
