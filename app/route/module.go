package route

import (
	"sync"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/contract"
)

// 使用contract包中的接口定义，保持向后兼容
type RouteModule = contract.RouteModule
type ModuleInitializer = contract.ModuleInitializer
type ModuleAutoRegister = contract.ModuleAutoRegister

// RouteModuleFunc 路由模块函数类型，用于简化模块注册
type RouteModuleFunc func(group *gin.RouterGroup, controllers interface{})

// RegisterRoutes 实现 RouteModule 接口
func (f RouteModuleFunc) RegisterRoutes(group *gin.RouterGroup, controllers interface{}) {
	f(group, controllers)
}

// Name 实现 RouteModule 接口
func (f RouteModuleFunc) Name() string {
	return "RouteModuleFunc"
}

// ModuleRegister 模块注册器，用于管理和注册所有路由模块
type ModuleRegister struct {
	modules []RouteModule
	mu      sync.RWMutex
}

// NewModuleRegister 创建一个新的模块注册器
func NewModuleRegister() *ModuleRegister {
	return &ModuleRegister{
		modules: make([]RouteModule, 0),
	}
}

// RegisterModule 注册一个路由模块
func (r *ModuleRegister) RegisterModule(module RouteModule) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.modules = append(r.modules, module)
}

// ModuleEntry 模块条目，包含模块实例和初始化信息
type ModuleEntry struct {
	Module      RouteModule
	Initializer ModuleInitializer
}

// moduleRegistry 模块注册表，用于存储模块名称和对应的模块实例
var (
	moduleRegistry = make(map[string]*ModuleEntry)
	registryMutex  sync.RWMutex
)

// GetModuleRegistrySize 返回模块注册表的大小（用于测试）
func GetModuleRegistrySize() int {
	registryMutex.RLock()
	defer registryMutex.RUnlock()
	return len(moduleRegistry)
}

// GlobalRouteRegistry 全局路由注册表
var GlobalRouteRegistry = NewModuleRegister()

// RegisterGlobalModule 注册全局路由模块
func RegisterGlobalModule(module RouteModule) {
	GlobalRouteRegistry.RegisterModule(module)
}

// RegisterModuleByName 注册模块到注册表
// name: 模块名称
// module: 模块实例
func RegisterModuleByName(name string, module RouteModule) {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	entry := &ModuleEntry{Module: module}
	if initializer, ok := module.(ModuleInitializer); ok {
		entry.Initializer = initializer
	}
	moduleRegistry[name] = entry
}

// AutoRegisterModule 自动注册模块路由
// name: 模块名称
// group: 路由组
// 返回: 是否成功注册
func AutoRegisterModule(name string, group *gin.RouterGroup) bool {
	fmt.Printf("AutoRegister module: %s\n", name)
	registryMutex.RLock()
	entry, exists := moduleRegistry[name]
	registryMutex.RUnlock()

	if !exists {
		return false
	}

	// 如果模块有初始化器，先初始化模块
	if entry.Initializer != nil && entry.Module == nil {
		entry.Module = entry.Initializer.Init()
	}

	// 注册模块路由
	entry.Module.RegisterRoutes(group, entry.Module)

	// 注册完成后移除对应的注册表项
	registryMutex.Lock()
	delete(moduleRegistry, name)
	registryMutex.Unlock()
	fmt.Printf("Removed module from registry: %s\n", name)

	return true
}

// AutoRegisterAllModules 自动注册所有已注册的模块路由
// group: 路由组
func AutoRegisterAllModules(group *gin.RouterGroup) {
	// 先获取所有模块名称和条目
	registryMutex.RLock()
	moduleNames := make([]string, 0, len(moduleRegistry))
	entries := make([]*ModuleEntry, 0, len(moduleRegistry))
	for name, entry := range moduleRegistry {
		moduleNames = append(moduleNames, name)
		entries = append(entries, entry)
	}
	registryMutex.RUnlock()

	for _, entry := range entries {
		// 如果模块有初始化器，先初始化模块
		if entry.Initializer != nil && entry.Module == nil {
			entry.Module = entry.Initializer.Init()
		}
		// 注册模块路由
		entry.Module.RegisterRoutes(group, entry.Module)
	}

	// 注册完成后移除所有已处理的模块条目
	registryMutex.Lock()
	for _, name := range moduleNames {
		delete(moduleRegistry, name)
	}
	registryMutex.Unlock()
}
