package route

import (
	"github.com/gin-gonic/gin"
	wms_admin_controller "github.com/maxlcoder/homework-backend/app/modules/wms/admin/controller"
	wms_api_controller "github.com/maxlcoder/homework-backend/app/modules/wms/api/controller"
	repository "github.com/maxlcoder/homework-backend/app/modules/wms/repository"
	wms_service "github.com/maxlcoder/homework-backend/app/modules/wms/service"
	"github.com/maxlcoder/homework-backend/database"

	"gorm.io/gorm"
)

type AdminController struct {
	PickingCarController *wms_admin_controller.PickingCarController
	BinController        *wms_admin_controller.BinController
	StaffController      *wms_admin_controller.StaffController
}

type ApiController struct {
	BinController *wms_api_controller.BinController
}

// WmsModule WMS模块结构
type WmsModule struct {
	AdminController *AdminController
	ApiController   *ApiController
}

// InitWmsModule 初始化WMS模块（导出方法）
func InitWmsModule(db *gorm.DB) *WmsModule {
	// 初始化仓库
	pickingCarRepository := repository.NewPickingCarRepository(db)
	binRepository := repository.NewBinRepository(db)
	staffRepository := repository.NewStaffRepository(db)

	// 初始化服务
	pickingCarService := wms_service.NewPickingCarService(db, pickingCarRepository)
	binService := wms_service.NewBinService(db, binRepository)
	staffService := wms_service.NewStaffService(db, staffRepository)

	// 初始化控制器
	adminPickingCarController := wms_admin_controller.NewPickingCarController(pickingCarService)
	adminBinController := wms_admin_controller.NewBinController(binService)
	adminStaffController := wms_admin_controller.NewStaffController(staffService)
	binController := wms_api_controller.NewBinController(binService)

	return &WmsModule{
		AdminController: &AdminController{
			PickingCarController: adminPickingCarController,
			BinController:        adminBinController,
			StaffController:      adminStaffController,
		},
		ApiController: &ApiController{
			BinController: binController,
		},
	}
}

// WMSApiRouteModule WMS API模块路由
type WMSApiRouteModule struct{}

// Name 返回模块名称
func (m *WMSApiRouteModule) Name() string {
	return "WMSApiRouteModule"
}

// WMSAutoModule WMS自动注册模块，实现RouteModule和ModuleInitializer接口
type WMSAutoModule struct {
	Module *WmsModule
}

// Name 返回模块名称
func (m *WMSAutoModule) Name() string {
	return "WMSAutoModule"
}

// Init 初始化模块，实现ModuleInitializer接口
func (m *WMSAutoModule) Init() interface{} {
	if m.Module == nil {
		m.Module = InitWmsModule(database.DB)
	}
	return m.Module.AdminController
}

// RegisterRoutes 注册模块路由，实现RouteModule接口
func (m *WMSAutoModule) RegisterRoutes(group *gin.RouterGroup) {
	// 注册前台路由
	m.Module.ApiController.RegisterRoutes(group)
	// 注册管理路由
	m.Module.AdminController.RegisterRoutes(group)
}

// RegisterRoutes 为 ApiController 添加路由注册方法
func (ctrl *ApiController) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("bins", ctrl.BinController.Page) // 分页列表
}

// RegisterRoutes 为 AdminController 添加路由注册方法
func (ctrl *AdminController) RegisterRoutes(group *gin.RouterGroup) {
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
