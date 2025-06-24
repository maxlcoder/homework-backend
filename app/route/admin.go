package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 模块
func RegisterAdminRoute(r *gin.RouterGroup, ctrl *AdminControllers, handle *jwt.GinJWTMiddleware) {
	r.POST("admins:register", ctrl.AdminController.Register) // 注册
	r.POST("login", handle.LoginHandler)
}

func RegisterAdminAuthRoute(r *gin.RouterGroup, ctrl *AdminControllers) {
	r.GET("users", ctrl.UserController.Page)   // 用户列表
	r.GET("admins", ctrl.AdminController.Page) // 管理员列表
}
