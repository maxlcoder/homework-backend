package route

import (
	"github.com/maxlcoder/homework-backend/app/controller"
)

type ApiControllers struct {
	UserController *controller.UserController
}

type AdminControllers struct {
	UserController  *controller.AdminUserController
	AdminController *controller.AdminController
}
