package contract

import (
	"sync"

	"github.com/maxlcoder/homework-backend/model"
)

// MenuProvider 菜单提供者接口（与 route_interface.go 保持一致）
type MenuProvider interface {
	// GetMenus 返回模块的菜单定义
	GetMenus() []model.Menu
}

// MenuRegistrar 菜单注册器接口，用于注册和管理模块菜单
type MenuRegister interface {
	// RegisterMenu 注册模块菜单
	RegisterMenu(provider MenuProvider)
	// GetAllMenus 获取所有模块的菜单定义
	GetAllMenus() []model.Menu
}

// BaseMenuRegistrar 基础菜单注册器实现，使用与路由注册相同的模式
type BaseMenuRegister struct {
	menuProviders []MenuProvider
	mu            sync.RWMutex
}

// NewBaseMenuRegistrar 创建新的菜单注册器
func NewBaseMenuRegister() *BaseMenuRegister {
	return &BaseMenuRegister{
		menuProviders: make([]MenuProvider, 0),
	}
}

// RegisterMenu 注册模块菜单
func (r *BaseMenuRegister) RegisterMenu(provider MenuProvider) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.menuProviders = append(r.menuProviders, provider)
}

// GetAllMenus 获取所有模块的菜单定义
func (r *BaseMenuRegister) GetAllMenus() []model.Menu {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var allMenus []model.Menu
	for _, provider := range r.menuProviders {
		menus := provider.GetMenus()
		allMenus = append(allMenus, menus...)
	}
	return allMenus
}

// 菜单注册表，与路由注册表类似
var (
	menuRegistry = make(map[string]MenuProvider)
	menuMutex    sync.RWMutex
)

// GlobalMenuRegistry 全局菜单注册表
var GlobalMenuRegistry = NewBaseMenuRegister()

// RegisterGlobalMenu 注册全局菜单提供者
func RegisterGlobalMenu(provider MenuProvider) {
	GlobalMenuRegistry.RegisterMenu(provider)
}

// RegisterMenuProvider 注册菜单提供者到注册表
// name: 菜单提供者名称
// provider: 菜单提供者实例
func RegisterMenuProvider(name string, provider MenuProvider) {
	menuMutex.Lock()
	defer menuMutex.Unlock()
	menuRegistry[name] = provider
}

// GetAllMenuProviders 获取所有菜单提供者
func GetAllMenuProviders() []MenuProvider {
	menuMutex.RLock()
	defer menuMutex.RUnlock()

	providers := make([]MenuProvider, 0, len(menuRegistry))
	for _, provider := range menuRegistry {
		providers = append(providers, provider)
	}
	return providers
}

// GetMenuProviderByName 按名称获取菜单提供者
func GetMenuProviderByName(name string) MenuProvider {
	menuMutex.RLock()
	defer menuMutex.RUnlock()
	return menuRegistry[name]
}

// GetAllMenus 获取所有模块的菜单定义（便捷方法）
func GetAllMenus() []model.Menu {
	var allMenus []model.Menu

	// 获取所有已注册的菜单提供者
	providers := GetAllMenuProviders()
	for _, provider := range providers {
		menus := provider.GetMenus()
		allMenus = append(allMenus, menus...)
	}

	// 同时包含全局菜单注册表中的菜单
	allMenus = append(allMenus, GlobalMenuRegistry.GetAllMenus()...)

	return allMenus
}
