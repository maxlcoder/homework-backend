package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 模块
func RegisterUserRoute(r *gin.RouterGroup, ctrl *ApiControllers, handle *jwt.GinJWTMiddleware) {
	r.POST("users:register", ctrl.UserController.Register) // 注册
	r.POST("login", handle.LoginHandler)                   // 登录
}

func RegisterUserAuthRoute(r *gin.RouterGroup, ctrl *ApiControllers) {
	r.GET("me", ctrl.UserController.Me) // 个人信息
}
