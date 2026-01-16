package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/contract"
	core_model "github.com/maxlcoder/homework-backend/app/modules/core/model"
	wms_admin_controller "github.com/maxlcoder/homework-backend/app/modules/wms/admin/controller"
	admin_middleware "github.com/maxlcoder/homework-backend/app/modules/wms/admin/middleware"
	wms_api_controller "github.com/maxlcoder/homework-backend/app/modules/wms/api/controller"
	api_middleware "github.com/maxlcoder/homework-backend/app/modules/wms/api/middleware"
	module_middleware "github.com/maxlcoder/homework-backend/app/modules/wms/middleware"
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	"github.com/maxlcoder/homework-backend/app/modules/wms/repository"
	"github.com/maxlcoder/homework-backend/app/modules/wms/service"

	"gorm.io/gorm"
)

// AdminController WMS管理后台控制器结构
type AdminController struct {
	PickingCarController    *wms_admin_controller.PickingCarController
	PickingBasketController *wms_admin_controller.PickingBasketController
	BinController           *wms_admin_controller.BinController
	StaffController         *wms_admin_controller.StaffController
}

// ApiController WMS API控制器结构
type ApiController struct {
	BinController *wms_api_controller.BinController
}

// WmsModule WMS模块结构，实现RouteModule和ModuleInitializer接口
type WmsModule struct {
	DB              *gorm.DB
	AdminController *AdminController
	ApiController   *ApiController
	initialized     bool
}

// Name 返回模块名称，实现RouteModule接口
func (m *WmsModule) Name() string {
	return "WmsModule"
}

// GetMenus 返回WMS模块的菜单定义，实现MenuProvider接口
func (m *WmsModule) GetMenus() []core_model.Menu {
	return []core_model.Menu{
		{
			Number: "wms-management",
			Name:   "WMS管理",
			Sort:   2,
			Children: []*core_model.Menu{
				{
					Number: "bin-management",
					Name:   "库位管理",
					Children: []*core_model.Menu{
						{
							Number: "bin-list",
							Name:   "列表",
							Permissions: []*core_model.Permission{
								{
									Name:   "列表",
									PATH:   "/admin/wms/bins",
									Method: "GET",
								},
							},
						},
						{
							Number: "bin-add",
							Name:   "新增",
							Permissions: []*core_model.Permission{
								{
									Name:   "新增",
									PATH:   "/admin/wms/bins",
									Method: "POST",
								},
							},
						},
						{
							Number: "bin-update",
							Name:   "更新",
							Permissions: []*core_model.Permission{
								{
									Name:   "更新",
									PATH:   "/admin/wms/bins/:id",
									Method: "PUT",
								},
							},
						},
						{
							Number: "bin-detail",
							Name:   "详情",
							Permissions: []*core_model.Permission{
								{
									Name:   "详情",
									PATH:   "/admin/wms/bins/:id",
									Method: "GET",
								},
							},
						},
						{
							Number: "bin-delete",
							Name:   "删除",
							Permissions: []*core_model.Permission{
								{
									Name:   "删除",
									PATH:   "/admin/wms/bins/:id",
									Method: "DELETE",
								},
							},
						},
					},
				},
				{
					Number: "picking-car-management",
					Name:   "拣货车辆管理",
					Children: []*core_model.Menu{
						{
							Number: "picking-car-list",
							Name:   "列表",
							Permissions: []*core_model.Permission{
								{
									Name:   "列表",
									PATH:   "/admin/wms/picking-cars",
									Method: "GET",
								},
							},
						},
						{
							Number: "picking-car-add",
							Name:   "新增",
							Permissions: []*core_model.Permission{
								{
									Name:   "新增",
									PATH:   "/admin/wms/picking-cars",
									Method: "POST",
								},
							},
						},
						{
							Number: "picking-car-update",
							Name:   "更新",
							Permissions: []*core_model.Permission{
								{
									Name:   "更新",
									PATH:   "/admin/wms/picking-cars/:id",
									Method: "PUT",
								},
							},
						},
						{
							Number: "picking-car-detail",
							Name:   "详情",
							Permissions: []*core_model.Permission{
								{
									Name:   "详情",
									PATH:   "/admin/wms/picking-cars/:id",
									Method: "GET",
								},
							},
						},
						{
							Number: "picking-car-delete",
							Name:   "删除",
							Permissions: []*core_model.Permission{
								{
									Name:   "删除",
									PATH:   "/admin/wms/picking-cars/:id",
									Method: "DELETE",
								},
							},
						},
					},
				},
				{
					Number: "staff-management",
					Name:   "员工管理",
					Children: []*core_model.Menu{
						{
							Number: "staff-list",
							Name:   "列表",
							Permissions: []*core_model.Permission{
								{
									Name:   "列表",
									PATH:   "/admin/wms/staff",
									Method: "GET",
								},
							},
						},
						{
							Number: "staff-add",
							Name:   "新增",
							Permissions: []*core_model.Permission{
								{
									Name:   "新增",
									PATH:   "/admin/wms/staff",
									Method: "POST",
								},
							},
						},
						{
							Number: "staff-update",
							Name:   "更新",
							Permissions: []*core_model.Permission{
								{
									Name:   "更新",
									PATH:   "/admin/wms/staff/:id",
									Method: "PUT",
								},
							},
						},
						{
							Number: "staff-detail",
							Name:   "详情",
							Permissions: []*core_model.Permission{
								{
									Name:   "详情",
									PATH:   "/admin/wms/staff/:id",
									Method: "GET",
								},
							},
						},
						{
							Number: "staff-delete",
							Name:   "删除",
							Permissions: []*core_model.Permission{
								{
									Name:   "删除",
									PATH:   "/admin/wms/staff/:id",
									Method: "DELETE",
								},
							},
						},
					},
				},
				{
					Number: "picking-basket-management",
					Name:   "拣货篮管理",
					Children: []*core_model.Menu{
						{
							Number: "picking-basket-list",
							Name:   "列表",
							Permissions: []*core_model.Permission{
								{
									Name:   "列表",
									PATH:   "/admin/wms/picking-baskets",
									Method: "GET",
								},
							},
						},
						{
							Number: "picking-basket-add",
							Name:   "新增",
							Permissions: []*core_model.Permission{
								{
									Name:   "新增",
									PATH:   "/admin/wms/picking-baskets",
									Method: "POST",
								},
							},
						},
						{
							Number: "picking-basket-update",
							Name:   "更新",
							Permissions: []*core_model.Permission{
								{
									Name:   "更新",
									PATH:   "/admin/wms/picking-baskets/:id",
									Method: "PUT",
								},
							},
						},
						{
							Number: "picking-basket-detail",
							Name:   "详情",
							Permissions: []*core_model.Permission{
								{
									Name:   "详情",
									PATH:   "/admin/wms/picking-baskets/:id",
									Method: "GET",
								},
							},
						},
						{
							Number: "picking-basket-delete",
							Name:   "删除",
							Permissions: []*core_model.Permission{
								{
									Name:   "删除",
									PATH:   "/admin/wms/picking-baskets/:id",
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
func (m *WmsModule) Init() contract.Module {
	if !m.initialized {
		// 初始化表
		model.AutoMigrate(m.DB)

		// 初始化仓库
		pickingCarRepository := repository.NewPickingCarRepository(m.DB)
		pickingBasketRepository := repository.NewPickingBasketRepository(m.DB)
		binRepository := repository.NewBinRepository(m.DB)
		staffRepository := repository.NewStaffRepository(m.DB)

		// 初始化服务
		pickingCarService := service.NewPickingCarService(m.DB, pickingCarRepository)
		pickingBasketService := service.NewPickingBasketService(m.DB, pickingBasketRepository)
		binService := service.NewBinService(m.DB, binRepository)
		staffService := service.NewStaffService(m.DB, staffRepository)

		// 初始化控制器
		adminPickingCarController := wms_admin_controller.NewPickingCarController(pickingCarService)
		adminPickingBasketController := wms_admin_controller.NewPickingBasketController(pickingBasketService)
		adminBinController := wms_admin_controller.NewBinController(binService)
		adminStaffController := wms_admin_controller.NewStaffController(staffService)
		binController := wms_api_controller.NewBinController(binService)

		// 设置控制器
		m.AdminController = &AdminController{
			PickingCarController:    adminPickingCarController,
			PickingBasketController: adminPickingBasketController,
			BinController:           adminBinController,
			StaffController:         adminStaffController,
		}
		m.ApiController = &ApiController{
			BinController: binController,
		}
		m.initialized = true
	}
	return m
}

// RegisterRoutes 注册模块路由，实现RouteModule接口
func (m *WmsModule) RegisterRoutes(apiGroup *gin.RouterGroup, apiAuthGroup *gin.RouterGroup, adminGroup *gin.RouterGroup, adminAuthGroup *gin.RouterGroup, module interface{}) {
	fmt.Println("Registering WMS Module routes")

	// 确保模块已初始化
	m.Init()

	// 注册模块接口
	apiGroup = apiGroup.Group("/wms")
	apiAuthGroup = apiAuthGroup.Group("/wms")
	// 添加WMS API模块级中间件
	apiGroup.Use(module_middleware.Logger())
	apiAuthGroup.Use(module_middleware.Logger())
	if m.ApiController != nil {
		m.ApiController.RegisterRoutes(apiGroup, apiAuthGroup)
	}

	// 注册Admin路由 - 后台接口
	adminGroup = adminGroup.Group("/wms")
	adminAuthGroup = adminAuthGroup.Group("/wms")
	// 应用Admin子模块的中间件
	adminGroup.Use(module_middleware.Logger())
	adminAuthGroup.Use(module_middleware.Logger())
	if m.AdminController != nil {

		// 注册需要认证的路由
		m.AdminController.RegisterRoutes(adminGroup, adminAuthGroup)
	}
}

// NewWmsModule 创建一个新的WMS模块实例（导出方法）
func NewWmsModule(db *gorm.DB) *WmsModule {
	return &WmsModule{DB: db}
}

// RegisterRoutes 为 ApiController 添加路由注册方法
func (ctrl *ApiController) RegisterRoutes(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	// 注册中间件
	authGroup.Use(api_middleware.Logger())

	// 普通接口 - 继承父路由组的中间件
	authGroup.GET("bins", ctrl.BinController.Page) // 分页列表、
}

// RegisterRoutes 为 AdminController 添加路由注册方法
func (ctrl *AdminController) RegisterRoutes(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	// 注册中间件
	authGroup.Use(admin_middleware.Logger())

	// ------------ 拣货车管理 ------------
	authGroup.GET("picking-cars", ctrl.PickingCarController.Page)           // 分页列表
	authGroup.GET("picking-cars/:id", ctrl.PickingCarController.Show)       // 详情
	authGroup.POST("picking-cars", ctrl.PickingCarController.Store)         // 新增
	authGroup.PUT("picking-cars/:id", ctrl.PickingCarController.Update)     // 更新
	authGroup.DELETE("picking-cars/:id", ctrl.PickingCarController.Destroy) // 删除

	// ------------ 拣货篮管理 ------------
	authGroup.GET("picking-baskets", ctrl.PickingBasketController.Page)           // 分页列表
	authGroup.GET("picking-baskets/:id", ctrl.PickingBasketController.Show)       // 详情
	authGroup.POST("picking-baskets", ctrl.PickingBasketController.Store)         // 新增
	authGroup.PUT("picking-baskets/:id", ctrl.PickingBasketController.Update)     // 更新
	authGroup.DELETE("picking-baskets/:id", ctrl.PickingBasketController.Destroy) // 删除

	// ------------ 库位管理 ------------
	authGroup.GET("bins", ctrl.BinController.Page)           // 分页列表
	authGroup.GET("bins/:id", ctrl.BinController.Show)       // 详情
	authGroup.POST("bins", ctrl.BinController.Store)         // 新增
	authGroup.PUT("bins/:id", ctrl.BinController.Update)     // 更新
	authGroup.DELETE("bins/:id", ctrl.BinController.Destroy) // 删除

	// ------------ 仓库人员管理 ------------
	authGroup.GET("staffs", ctrl.StaffController.Page)           // 分页列表
	authGroup.GET("staffs/:id", ctrl.StaffController.Show)       // 详情
	authGroup.POST("staffs", ctrl.StaffController.Store)         // 新增
	authGroup.PUT("staffs/:id", ctrl.StaffController.Update)     // 更新
	authGroup.DELETE("staffs/:id", ctrl.StaffController.Destroy) // 删除
}
