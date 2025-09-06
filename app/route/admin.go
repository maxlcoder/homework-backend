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

	r.GET("roles", ctrl.RoleController.Page)       // 角色分页列表
	r.POST("roles", ctrl.RoleController.Store)     // 角色分页列表
	r.PUT("roles/:id", ctrl.RoleController.Update) // 角色分页列表
	r.DELETE("roles", ctrl.RoleController.Destroy) // 角色分页列表
	r.GET("roles/:id", ctrl.RoleController.Show)   // 角色分页列表

}
