package route

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// 简单的Admin模块初始化函数
func initAdminModule(group *gin.RouterGroup, controllers interface{}) {
	// 实现简单的管理员路由
	group.GET("/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Admin module"})
	})
}

func ApiRoutes(r *gin.Engine, enforcer *casbin.Enforcer) {
	// 1. 注册全局模块
	RegisterGlobalModuleFunc("AdminModule", initAdminModule)

	// 2. 创建API路由组
	api := r.Group("/api")

	// 3. 创建Admin路由组
	admin := r.Group("/admin")

	// 4. 自动注册所有全局模块路由
	GlobalRouteRegistry.RegisterAllRoutes(admin, nil)

	// 5. 示例：注册一个简单的API路由
	api.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from API"})
	})
}
