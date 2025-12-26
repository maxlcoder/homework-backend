package route

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	wms_route "github.com/maxlcoder/homework-backend/app/modules/wms/route"
	"github.com/maxlcoder/homework-backend/database"
)

// ApiRoutes 注册所有API路由
func ApiRoutes(r *gin.Engine, enforcer *casbin.Enforcer) {
	// 注册模块
	RegisterModuleByName("WmsModule", &wms_route.WmsModule{DB: database.DB})

	// 自动注册所有模块路由
	AutoRegisterAllModules(r.Group(""))
}
