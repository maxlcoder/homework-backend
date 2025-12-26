package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

// 用户模块初始化
func initUserModule(db *gorm.DB) *UserModule {
	userRepository := repository.NewUserRepository(db)
	userService := admin_service.NewUserService(userRepository) // 初始化控制器
	return &UserModule{
		UserController: admin_controller.NewUserController(userService),
	}
}

// 管理员模块初始化
func initAdminModule(db *gorm.DB, enforcer *casbin.Enforcer) *AdminModule {
	// 初始化仓库
	adminRepository := repository.NewAdminRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	adminRoleRepository := repository.NewAdminRoleRepository(db)
	menuRepository := repository.NewMenuRepository(db)
	roleMenuRepository := repository.NewRoleMenuRepository(db)

	// 初始化服务
	adminService := admin_service.NewAdminService(
		db,
		adminRepository,
		roleRepository,
		adminRoleRepository)
	roleService := admin_service.NewRoleService(
		db,
		enforcer,
		roleRepository,
		menuRepository,
		roleMenuRepository)
	userService := admin_service.NewUserService(repository.NewUserRepository(db))

	// 初始化控制器
	return &AdminModule{
		UserController:  admin_controller.NewAdminUserController(adminService, userService),
		AdminController: admin_controller.NewAdminController(adminService),
		RoleController:  admin_controller.NewRoleController(roleService),
	}
}

// APIAuthRouteModule API认证路由模块
type APIAuthRouteModule struct{}

// Name 返回模块名称
func (m *APIAuthRouteModule) Name() string {
	return "APIAuthRouteModule"
}

// RegisterRoutes 注册API认证路由
func (m *APIAuthRouteModule) RegisterRoutes(group *gin.RouterGroup, controllers interface{}) {
	ctrl, ok := controllers.(*ApiControllers)
	if !ok {
		return
	}

	// 注册认证后才能访问的路由
	group.GET("me", ctrl.UserController.Me) // 个人信息
}

// AdminRouteModule 管理员路由模块
type AdminRouteModule struct {
	AuthMiddleware *jwt.GinJWTMiddleware
}

// Name 返回模块名称
func (m *AdminRouteModule) Name() string {
	return "AdminRouteModule"
}

// RegisterRoutes 注册管理员路由
func (m *AdminRouteModule) RegisterRoutes(group *gin.RouterGroup, controllers interface{}) {
	ctrl, ok := controllers.(*AdminControllers)
	if !ok {
		return
	}

	// 注册管理员相关路由
	group.POST("admins:register", ctrl.AdminController.Register) // 注册
	group.POST("login", m.AuthMiddleware.LoginHandler)
}

// AdminAuthRouteModule 管理员认证路由模块
type AdminAuthRouteModule struct{}

// Name 返回模块名称
func (m *AdminAuthRouteModule) Name() string {
	return "AdminAuthRouteModule"
}

// RegisterRoutes 注册管理员认证路由
func (m *AdminAuthRouteModule) RegisterRoutes(group *gin.RouterGroup, controllers interface{}) {

	// ---------- 业务功能 ----------

	group.GET("users", ctrl.UserController.Page) // 用户列表

	// ---------- 平台功能 ----------
	// ------------ 个人中心 ------------
	group.GET("me", ctrl.AdminController.Me)

	// ------------ 管理员管理 ------------
	group.GET("admins", ctrl.AdminController.Page)           // 分页列表
	group.GET("admins/:id", ctrl.AdminController.Show)       // 详情
	group.POST("admins", ctrl.AdminController.Store)         // 新增
	group.PUT("admins/:id", ctrl.AdminController.Update)     // 更新
	group.DELETE("admins/:id", ctrl.AdminController.Destroy) // 删除

	// ------------ 角色管理 ------------
	group.GET("roles", ctrl.RoleController.Page)           // 分页列表
	group.GET("roles/:id", ctrl.RoleController.Show)       // 详情
	group.POST("roles", ctrl.RoleController.Store)         // 新增
	group.PUT("roles/:id", ctrl.RoleController.Update)     // 更新
	group.DELETE("roles/:id", ctrl.RoleController.Destroy) // 删除
}
