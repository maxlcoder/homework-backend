package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/controller"
	"github.com/maxlcoder/homework-backend/service"
)

// 模块
func RegisterUserRoute(r *gin.RouterGroup, ctrl *ApiController, handle *jwt.GinJWTMiddleware) {
	// 初始化控制器，注入服务
	r.POST("users:register", ctrl.UserController.Register) // 注册
	//r.POST("login", ctrl.UserController.Login)             // 登录
	r.POST("login", handle.LoginHandler) // 登录
}

func RegisterUserAuthRoute(r *gin.RouterGroup, ctrl *ApiController) {
	r.GET("me", ctrl.UserController.Me) // 个人信息
}

func AdminRegisterUserRoute(r *gin.RouterGroup) {
	userService := &service.UserService{}
	userController := controller.NewUserController(userService)
	r.GET("users", userController.Register)
}
