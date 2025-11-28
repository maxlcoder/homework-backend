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
	// ---------- 业务功能 ----------

	r.GET("users", ctrl.UserController.Page) // 用户列表

	// ---------- 平台功能 ----------
	// ------------ 个人中心 ------------
	r.GET("me", ctrl.AdminController.Me)

	// ------------ 管理员管理 ------------
	r.GET("admins", ctrl.AdminController.Page)           // 分页列表
	r.POST("admins", ctrl.AdminController.Store)         // 新增
	r.PUT("admins/:id", ctrl.AdminController.Update)     // 更新
	r.DELETE("admins/:id", ctrl.AdminController.Destroy) // 删除
	r.GET("admins/:id", ctrl.AdminController.Show)       // 详情

	// ------------ 角色管理 ------------
	r.GET("roles", ctrl.RoleController.Page)           // 分页列表
	r.POST("roles", ctrl.RoleController.Store)         // 新增
	r.PUT("roles/:id", ctrl.RoleController.Update)     // 更新
	r.DELETE("roles/:id", ctrl.RoleController.Destroy) // 删除
	r.GET("roles/:id", ctrl.RoleController.Show)       // 详情

}
