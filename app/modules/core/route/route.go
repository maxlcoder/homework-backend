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
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	core_model "github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/app/modules/core/service"
	"gorm.io/gorm"
)

type AdminController struct {
	UserController   *admin_controller.AdminUserController
	AdminController  *admin_controller.AdminController
	RoleController   *admin_controller.RoleController
	TenantController *admin_controller.TenantController
	Handler          *jwt.GinJWTMiddleware
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

// GetMenus 返回核心模块的菜单定义，实现MenuProvider接口
func (m *CoreModule) GetMenus() []core_model.Menu {
	return []core_model.Menu{
		{
			Number: "system-setting",
			Name:   "系统设置",
			Sort:   1,
			Children: []*core_model.Menu{
				{
					Number: "admin-management",
					Name:   "账号管理",
					Children: []*core_model.Menu{
						{
							Number: "admin-list",
							Name:   "列表",
							Permissions: []*core_model.Permission{
								{
									Name:   "列表",
									PATH:   "/admin/admins",
									Method: "GET",
								},
							},
						},
						{
							Number: "admin-add",
							Name:   "新增",
							Permissions: []*core_model.Permission{
								{
									Name:   "新增",
									PATH:   "/admin/admins",
									Method: "POST",
								},
							},
						},
						{
							Number: "admin-update",
							Name:   "更新",
							Permissions: []*core_model.Permission{
								{
									Name:   "更新",
									PATH:   "/admin/admins/:id",
									Method: "PUT",
								},
							},
						},
						{
							Number: "admin-detail",
							Name:   "详情",
							Permissions: []*core_model.Permission{
								{
									Name:   "详情",
									PATH:   "/admin/admins/:id",
									Method: "GET",
								},
							},
						},
						{
							Number: "admin-delete",
							Name:   "删除",
							Permissions: []*core_model.Permission{
								{
									Name:   "删除",
									PATH:   "/admin/admins/:id",
									Method: "DELETE",
								},
							},
						},
					},
				},
				{
					Number: "role-management",
					Name:   "角色管理",
					Children: []*core_model.Menu{
						{
							Number: "role-list",
							Name:   "列表",
							Permissions: []*core_model.Permission{
								{
									Name:   "列表",
									PATH:   "/admin/roles",
									Method: "GET",
								},
							},
						},
						{
							Number: "role-add",
							Name:   "新增",
							Permissions: []*core_model.Permission{
								{
									Name:   "新增",
									PATH:   "/admin/roles",
									Method: "POST",
								},
							},
						},
						{
							Number: "role-update",
							Name:   "更新",
							Permissions: []*core_model.Permission{
								{
									Name:   "更新",
									PATH:   "/admin/roles/:id",
									Method: "PUT",
								},
							},
						},
						{
							Number: "role-detail",
							Name:   "详情",
							Permissions: []*core_model.Permission{
								{
									Name:   "详情",
									PATH:   "/admin/roles/:id",
									Method: "GET",
								},
							},
						},
						{
							Number: "role-delete",
							Name:   "删除",
							Permissions: []*core_model.Permission{
								{
									Name:   "删除",
									PATH:   "/admin/roles/:id",
									Method: "DELETE",
								},
							},
						},
					},
				},
				{
					Number: "tenant-management",
					Name:   "租户管理",
					Children: []*core_model.Menu{
						{
							Number: "tenant-list",
							Name:   "列表",
							Permissions: []*core_model.Permission{
								{
									Name:   "列表",
									PATH:   "/admin/tenants",
									Method: "GET",
								},
							},
						},
						{
							Number: "tenant-add",
							Name:   "新增",
							Permissions: []*core_model.Permission{
								{
									Name:   "新增",
									PATH:   "/admin/tenants",
									Method: "POST",
								},
							},
						},
						{
							Number: "tenant-update",
							Name:   "更新",
							Permissions: []*core_model.Permission{
								{
									Name:   "更新",
									PATH:   "/admin/tenants/:id",
									Method: "PUT",
								},
							},
						},
						{
							Number: "tenant-detail",
							Name:   "详情",
							Permissions: []*core_model.Permission{
								{
									Name:   "详情",
									PATH:   "/admin/tenants/:id",
									Method: "GET",
								},
							},
						},
						{
							Number: "tenant-delete",
							Name:   "删除",
							Permissions: []*core_model.Permission{
								{
									Name:   "删除",
									PATH:   "/admin/tenants/:id",
									Method: "DELETE",
								},
							},
						},
					},
				},
			},
		},
	}
}

// Init 初始化模块，实现ModuleInitializer接口
func (m *CoreModule) Init() contract.Module {
	if !m.initialized {

		// 初始化表
		model.AutoMigrate(m.DB)

		// 初始化仓库

		userService := service.NewUserService(m.DB) // 初始化控制器
		menuService := service.NewMenuService(m.DB)
		// 初始化服务
		adminService := service.NewAdminService(m.DB)
		roleService := service.NewRoleService(m.DB, m.Enforcer, menuService)
		tenantService := service.NewTenantService(m.DB)

		m.ApiController = &ApiController{
			UserController: api_controller.NewUserController(userService),
		}
		m.AdminController = &AdminController{
			UserController:   admin_controller.NewAdminUserController(adminService, userService),
			AdminController:  admin_controller.NewAdminController(adminService),
			RoleController:   admin_controller.NewRoleController(roleService),
			TenantController: admin_controller.NewTenantController(tenantService),
			Handler:          m.AdminHandler,
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

	// ------------ 租户管理 ------------
	authGroup.GET("tenants", ctrl.TenantController.Page)           // 分页列表
	authGroup.GET("tenants/:id", ctrl.TenantController.Show)       // 详情
	authGroup.POST("tenants", ctrl.TenantController.Store)         // 新增
	authGroup.PUT("tenants/:id", ctrl.TenantController.Update)     // 更新
	authGroup.DELETE("tenants/:id", ctrl.TenantController.Destroy) // 删除
}
