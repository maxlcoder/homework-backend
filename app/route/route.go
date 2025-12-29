package route

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/middleware"
	wms_route "github.com/maxlcoder/homework-backend/app/modules/wms/route"
	"github.com/maxlcoder/homework-backend/app/route/auth"
	"github.com/maxlcoder/homework-backend/database"
)

// ApiRoutes 注册所有API路由
func ApiRoutes(r *gin.Engine, enforcer *casbin.Enforcer) {
	// 全局公用中间件 - 应用于所有路由
	// CORS中间件
	r.Use(middleware.Cors())
	// 错误处理中间件
	r.Use(middleware.ErrorHandler())
	// 请求日志中间件
	r.Use(middleware.Logger())

	// 注册WMS模块 - 它可以在自己的Middleware方法中定义特定的中间件
	RegisterModuleByName("WmsModule", &wms_route.WmsModule{DB: database.DB})

	// auth 中间件 - 可作为模块级别的公用中间件
	authMiddleware, err := jwt.New(auth.InitJwtParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	adminAuthMiddleware, err := jwt.New(auth.InitAdminJwtParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// 可以创建不同的路由组，应用不同的公用中间件
	// 例如：API路由组
	apiGroup := r.Group("/api")
	// 添加系统整体中间件 - API组
	apiGroup.Use(auth.HandlerMiddleware(authMiddleware))
	// 系统整体中间件 - 可以添加更多系统级中间件

	// 管理后台路由组
	adminGroup := r.Group("/admin")
	// 管理后台路由组可以应用管理员特定的中间件
	adminGroup.Use(auth.HandlerMiddleware(adminAuthMiddleware))
	// 系统整体中间件 - 管理后台组
	adminGroup.Use(middleware.CasbinMiddleware(enforcer))

	// 自动注册所有模块路由
	// 这里使用根路由组，模块会根据自己的Middleware方法应用特定的中间件
	AutoRegisterAllModules(apiGroup, adminGroup)
}
