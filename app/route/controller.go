package route

import (
	admin_controller "github.com/maxlcoder/homework-backend/app/modules/core/admin/controller"
	api_controller "github.com/maxlcoder/homework-backend/app/modules/core/api/controller"
	wms_admin_controller "github.com/maxlcoder/homework-backend/app/modules/wms/admin/controller"
)

type ApiControllers struct {
	UserController *api_controller.UserController
}

type AdminControllers struct {
	UserController  *admin_controller.AdminUserController
	AdminController *admin_controller.AdminController
	RoleController  *admin_controller.RoleController

	PickingCarController *wms_admin_controller.PickingCarController
	BinController        *wms_admin_controller.BinController
	StaffController      *wms_admin_controller.StaffController
}
