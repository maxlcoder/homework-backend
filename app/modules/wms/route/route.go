package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/contract"
	wms_admin_controller "github.com/maxlcoder/homework-backend/app/modules/wms/admin/controller"
	wms_api_controller "github.com/maxlcoder/homework-backend/app/modules/wms/api/controller"
	"github.com/maxlcoder/homework-backend/app/modules/wms/repository"
	"github.com/maxlcoder/homework-backend/app/modules/wms/service"

	admin_middleware "github.com/maxlcoder/homework-backend/app/modules/wms/admin/middleware"
	api_middleware "github.com/maxlcoder/homework-backend/app/modules/wms/api/middleware"
	module_middleware "github.com/maxlcoder/homework-backend/app/modules/wms/middleware"

	"gorm.io/gorm"
)

// AdminController WMS管理后台控制器结构
type AdminController struct {
	PickingCarController *wms_admin_controller.PickingCarController
	BinController        *wms_admin_controller.BinController
	StaffController      *wms_admin_controller.StaffController
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

// Init 初始化模块，实现ModuleInitializer接口
func (m *WmsModule) Init() contract.RouteModule {
	if !m.initialized {
		// 初始化仓库
		pickingCarRepository := repository.NewPickingCarRepository(m.DB)
		binRepository := repository.NewBinRepository(m.DB)
		staffRepository := repository.NewStaffRepository(m.DB)

		// 初始化服务
		pickingCarService := service.NewPickingCarService(m.DB, pickingCarRepository)
		binService := service.NewBinService(m.DB, binRepository)
		staffService := service.NewStaffService(m.DB, staffRepository)

		// 初始化控制器
		adminPickingCarController := wms_admin_controller.NewPickingCarController(pickingCarService)
		adminBinController := wms_admin_controller.NewBinController(binService)
		adminStaffController := wms_admin_controller.NewStaffController(staffService)
		binController := wms_api_controller.NewBinController(binService)

		// 设置控制器
		m.AdminController = &AdminController{
			PickingCarController: adminPickingCarController,
			BinController:        adminBinController,
			StaffController:      adminStaffController,
		}
		m.ApiController = &ApiController{
			BinController: binController,
		}
		m.initialized = true
	}
	return m
}

// RegisterRoutes 注册模块路由，实现RouteModule接口
func (m *WmsModule) RegisterRoutes(apiGroup *gin.RouterGroup, adminGroup *gin.RouterGroup, module interface{}) {
	fmt.Println("Registering WMS Module routes")

	// 确保模块已初始化
	m.Init()

	// 注册模块接口
	apiGroup = apiGroup.Group("/wms")
	// 添加WMS API模块级中间件
	apiGroup.Use(module_middleware.Logger())
	if m.ApiController != nil {
		m.ApiController.RegisterRoutes(apiGroup)
	}

	// 注册Admin路由 - 后台接口
	adminGroup = adminGroup.Group("/wms")
	// 应用Admin子模块的中间件
	adminGroup.Use(module_middleware.Logger())
	if m.AdminController != nil {
		m.AdminController.RegisterRoutes(adminGroup)
	}
}

// NewWmsModule 创建一个新的WMS模块实例（导出方法）
func NewWmsModule(db *gorm.DB) *WmsModule {
	return &WmsModule{DB: db}
}

// RegisterRoutes 为 ApiController 添加路由注册方法
func (ctrl *ApiController) RegisterRoutes(group *gin.RouterGroup) {
	// 注册中间件
	group.Use(api_middleware.Logger())

	// 普通接口 - 继承父路由组的中间件
	group.GET("bins", ctrl.BinController.Page) // 分页列表、
}

// RegisterRoutes 为 AdminController 添加路由注册方法
func (ctrl *AdminController) RegisterRoutes(group *gin.RouterGroup) {
	// 注册中间件
	group.Use(admin_middleware.Logger())

	// ------------ 拣货车管理 ------------
	group.GET("picking-cars", ctrl.PickingCarController.Page)           // 分页列表
	group.GET("picking-cars/:id", ctrl.PickingCarController.Show)       // 详情
	group.POST("picking-cars", ctrl.PickingCarController.Store)         // 新增
	group.PUT("picking-cars/:id", ctrl.PickingCarController.Update)     // 更新
	group.DELETE("picking-cars/:id", ctrl.PickingCarController.Destroy) // 删除

	// ------------ 库位管理 ------------
	group.GET("bins", ctrl.BinController.Page)           // 分页列表
	group.GET("bins/:id", ctrl.BinController.Show)       // 详情
	group.POST("bins", ctrl.BinController.Store)         // 新增
	group.PUT("bins/:id", ctrl.BinController.Update)     // 更新
	group.DELETE("bins/:id", ctrl.BinController.Destroy) // 删除

	// ------------ 仓库人员管理 ------------
	group.GET("staffs", ctrl.StaffController.Page)           // 分页列表
	group.GET("staffs/:id", ctrl.StaffController.Show)       // 详情
	group.POST("staffs", ctrl.StaffController.Store)         // 新增
	group.PUT("staffs/:id", ctrl.StaffController.Update)     // 更新
	group.DELETE("staffs/:id", ctrl.StaffController.Destroy) // 删除
}
