package route

import (
	"fmt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/contract"
	admin_controller "github.com/maxlcoder/homework-backend/app/modules/core/admin/controller"
	admin_middleware "github.com/maxlcoder/homework-backend/app/modules/core/admin/middleware"
	api_controller "github.com/maxlcoder/homework-backend/app/modules/core/api/controller"
	api_middleware "github.com/maxlcoder/homework-backend/app/modules/core/api/middleware"
	"github.com/maxlcoder/homework-backend/app/modules/core/service"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type AdminController struct {
	UserController  *admin_controller.AdminUserController
	AdminController *admin_controller.AdminController
	RoleController  *admin_controller.RoleController
	Handler         *jwt.GinJWTMiddleware
}

type ApiController struct {
	UserController *api_controller.UserController
}

type CoreModule struct {
	DB              *gorm.DB
	Enforcer        *casbin.Enforcer
	AdminController *AdminController
	ApiController   *ApiController
	initialized     bool
	ApiHandler      *jwt.GinJWTMiddleware
	AdminHandler    *jwt.GinJWTMiddleware
}

// Name 返回模块名称，实现RouteModule接口
func (m *CoreModule) Name() string {
	return "CoreModule"
}

// Init 初始化模块，实现ModuleInitializer接口
func (m *CoreModule) Init() contract.RouteModule {
	if !m.initialized {
		userRepository := repository.NewUserRepository(m.DB)
		// 初始化仓库
		adminRepository := repository.NewAdminRepository(m.DB)
		roleRepository := repository.NewRoleRepository(m.DB)
		adminRoleRepository := repository.NewAdminRoleRepository(m.DB)
		menuRepository := repository.NewMenuRepository(m.DB)
		roleMenuRepository := repository.NewRoleMenuRepository(m.DB)

		userService := service.NewUserService(userRepository) // 初始化控制器
		// 初始化服务
		adminService := service.NewAdminService(m.DB, adminRepository, roleRepository, adminRoleRepository)
		roleService := service.NewRoleService(m.DB, m.Enforcer, roleRepository, menuRepository, roleMenuRepository)

		m.ApiController = &ApiController{
			UserController: api_controller.NewUserController(userService),
		}
		m.AdminController = &AdminController{
			UserController:  admin_controller.NewAdminUserController(adminService, userService),
			AdminController: admin_controller.NewAdminController(adminService),
			RoleController:  admin_controller.NewRoleController(roleService),
			Handler:         m.AdminHandler,
		}
		m.initialized = true
	}
	return m
}

// RegisterRoutes 注册模块路由，实现RouteModule接口
func (m *CoreModule) RegisterRoutes(apiGroup *gin.RouterGroup, apiAuthGroup *gin.RouterGroup, adminGroup *gin.RouterGroup, adminAuthGroup *gin.RouterGroup, module interface{}) {
	fmt.Println("Registering Core Module routes")

	// 确保模块已初始化
	m.Init()

	// 添加 API模块级中间件
	apiGroup = apiGroup.Group("")
	apiGroup.Use(api_middleware.Logger())
	if m.ApiController != nil {
		m.ApiController.RegisterRoutes(apiGroup, apiAuthGroup)
	}

	// 注册Admin路由 - 后台接口
	adminGroup = adminGroup.Group("")
	// 应用Admin子模块的中间件
	adminGroup.Use(admin_middleware.Logger())
	if m.AdminController != nil {
		m.AdminController.RegisterRoutes(adminGroup, adminAuthGroup)
	}
}

// RegisterRoutes 注册API认证路由
func (ctrl *ApiController) RegisterRoutes(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	// 注册认证后才能访问的路由
	group.GET("me", ctrl.UserController.Me) // 个人信息
}

// RegisterRoutes 注册管理员认证路由
func (ctrl *AdminController) RegisterRoutes(group *gin.RouterGroup, authGroup *gin.RouterGroup) {

	// 注册管理员相关路由
	group.POST("admins:register", ctrl.AdminController.Register) // 注册
	group.POST("login", ctrl.Handler.LoginHandler)

	// ---------- 业务功能 ----------

	authGroup.GET("users", ctrl.UserController.Page) // 用户列表

	// ---------- 平台功能 ----------
	// ------------ 个人中心 ------------
	authGroup.GET("me", ctrl.AdminController.Me)

	// ------------ 管理员管理 ------------
	authGroup.GET("admins", ctrl.AdminController.Page)           // 分页列表
	authGroup.GET("admins/:id", ctrl.AdminController.Show)       // 详情
	authGroup.POST("admins", ctrl.AdminController.Store)         // 新增
	authGroup.PUT("admins/:id", ctrl.AdminController.Update)     // 更新
	authGroup.DELETE("admins/:id", ctrl.AdminController.Destroy) // 删除

	// ------------ 角色管理 ------------
	authGroup.GET("roles", ctrl.RoleController.Page)           // 分页列表
	authGroup.GET("roles/:id", ctrl.RoleController.Show)       // 详情
	authGroup.POST("roles", ctrl.RoleController.Store)         // 新增
	authGroup.PUT("roles/:id", ctrl.RoleController.Update)     // 更新
	authGroup.DELETE("roles/:id", ctrl.RoleController.Destroy) // 删除
}
