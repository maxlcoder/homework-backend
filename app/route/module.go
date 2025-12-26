package route

import (
	"sync"

	"github.com/gin-gonic/gin"
)

// ModuleInitializer 模块初始化器接口，用于自动初始化模块
type ModuleInitializer interface {
	// Init 初始化模块
	// 返回: 模块控制器
	Init() interface{}
}

// ModuleAutoRegister 模块自动注册接口，模块需要实现此接口才能被自动注册
type ModuleAutoRegister interface {
	RouteModule
	ModuleInitializer
}

// RouteModule 路由模块接口，每个模块需要实现此接口来注册自己的路由
type RouteModule interface {
	// RegisterRoutes 注册模块路由
	// group: 路由组
	// controllers: 控制器集合
	RegisterRoutes(group *gin.RouterGroup, controllers interface{})
	// Name 返回模块名称
	Name() string
}

// RouteModuleFunc 路由模块函数类型，用于简化模块注册
type RouteModuleFunc func(group *gin.RouterGroup, controllers interface{})

// RegisterRoutes 实现 RouteModule 接口
func (f RouteModuleFunc) RegisterRoutes(group *gin.RouterGroup, controllers interface{}) {
	f(group, controllers)
}

// Name 实现 RouteModule 接口
func (f RouteModuleFunc) Name() string {
	return "AnonymousModule"
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

// RegisterModuleFunc 注册一个路由模块函数
func (r *ModuleRegister) RegisterModuleFunc(name string, f RouteModuleFunc) {
	r.RegisterModule(routeModuleFuncWithName{name: name, f: f})
}

// routeModuleFuncWithName 带名称的路由模块函数
type routeModuleFuncWithName struct {
	name string
	f    RouteModuleFunc
}

// RegisterRoutes 实现 RouteModule 接口
func (r routeModuleFuncWithName) RegisterRoutes(group *gin.RouterGroup, controllers interface{}) {
	r.f(group, controllers)
}

// Name 实现 RouteModule 接口
func (r routeModuleFuncWithName) Name() string {
	return r.name
}

// RegisterAllRoutes 注册所有模块的路由
func (r *ModuleRegister) RegisterAllRoutes(group *gin.RouterGroup, controllers interface{}) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, module := range r.modules {
		module.RegisterRoutes(group, controllers)
	}
}

// ModuleEntry 模块条目，包含模块实例和初始化信息
type ModuleEntry struct {
	Module      RouteModule
	Initializer ModuleInitializer
	Controller  interface{}
}

// moduleRegistry 模块注册表，用于存储模块名称和对应的模块实例
var (
	moduleRegistry = make(map[string]*ModuleEntry)
	registryMutex  sync.RWMutex
)

// GlobalRouteRegistry 全局路由注册表
var GlobalRouteRegistry = NewModuleRegister()

// RegisterGlobalModule 注册全局路由模块
func RegisterGlobalModule(module RouteModule) {
	GlobalRouteRegistry.RegisterModule(module)
}

// RegisterGlobalModuleFunc 注册全局路由模块函数
func RegisterGlobalModuleFunc(name string, f RouteModuleFunc) {
	GlobalRouteRegistry.RegisterModuleFunc(name, f)
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
	registryMutex.RLock()
	entry, exists := moduleRegistry[name]
	registryMutex.RUnlock()

	if !exists {
		return false
	}

	// 如果模块有初始化器，先初始化模块
	if entry.Initializer != nil && entry.Controller == nil {
		entry.Controller = entry.Initializer.Init()
	}

	// 注册模块路由
	entry.Module.RegisterRoutes(group, entry.Controller)
	return true
}

// AutoRegisterAllModules 自动注册所有已注册的模块路由
// group: 路由组
func AutoRegisterAllModules(group *gin.RouterGroup) {
	registryMutex.RLock()
	entries := make([]*ModuleEntry, 0, len(moduleRegistry))
	for _, entry := range moduleRegistry {
		entries = append(entries, entry)
	}
	registryMutex.RUnlock()

	for _, entry := range entries {
		// 如果模块有初始化器，先初始化模块
		if entry.Initializer != nil && entry.Controller == nil {
			entry.Controller = entry.Initializer.Init()
		}

		// 注册模块路由
		entry.Module.RegisterRoutes(group, entry.Controller)
	}
}
