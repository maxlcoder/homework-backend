package contract

import (
	"github.com/gin-gonic/gin"
)

// 定义模块
type Module interface {
	// RegisterRoutes 注册模块路由
	// group: 路由组
	// controllers: 控制器集合
	RegisterRoutes(apiGroup *gin.RouterGroup, apiAuthGroup *gin.RouterGroup, adminGroup *gin.RouterGroup, adminAuthGroup *gin.RouterGroup, module interface{})
	// Name 返回模块名称
	Name() string
}

// ModuleInitializer 模块初始化器接口，用于自动初始化模块
type ModuleInitializer interface {
	// Init 初始化模块
	// 返回: 模块控制器
	Init() Module
}

// ModuleAutoRegister 模块自动注册接口，模块需要实现此接口才能被自动注册
type ModuleAutoRegister interface {
	Module
	ModuleInitializer
}
