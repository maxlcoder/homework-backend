package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	request "github.com/maxlcoder/homework-backend/app/request/user"
	"github.com/maxlcoder/homework-backend/app/response"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/validator"
	"github.com/maxlcoder/homework-backend/service"
)

type UserController struct {
	BaseController
	// 集成服务
	userService service.UserServiceInterface
}

// 初始化 controller，并注入服务
func NewUserController(userService service.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) Register(c *gin.Context) {
	var userCreateRequest request.UserCreateRequest

	// 进行参数校验
	if errs, ok := validator.BindAndValidateFirst(c, &userCreateRequest); !ok {
		controller.Error(c, 400, errs)
		return
	}

	var user model.User
	err := copier.Copy(&user, &userCreateRequest)
	if err != nil {
		controller.Error(c, 400, "数据获取失败")
		return
	}
	_, err = controller.userService.Create(&user)
	if err != nil {
		controller.Error(c, 400, fmt.Errorf("注册失败：%w", err).Error())
		return
	}
	dataID := response.DataId{ID: int(user.ID)}
	controller.Success(c, dataID)
}

func (controller *UserController) Login(c *gin.Context) {
	
}

func (controller *UserController) Me(c *gin.Context) {

}
