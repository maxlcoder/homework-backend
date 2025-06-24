package route

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/controller"
	"github.com/maxlcoder/homework-backend/app/route/auth"
	"github.com/maxlcoder/homework-backend/pkg/database"
	"github.com/maxlcoder/homework-backend/repository"
	"github.com/maxlcoder/homework-backend/service"
)

// service 列表
var (
	userRepository  repository.UserRepository
	adminRepository repository.AdminRepository
)

// service 列表
var (
	userService  *service.UserService
	adminService *service.AdminService
)

var (
	apiController   *ApiControllers
	adminController *AdminControllers
)

func ApiRoutes(r *gin.Engine) {

	// 注册中间件
	initRepository()
	initService()
	initController()

	// auth 中间件
	authMiddleware, err := jwt.New(auth.InitJwtParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	// 初始化
	r.Use(auth.HandlerMiddleware(authMiddleware))

	adminAuthMiddleware, err := jwt.New(auth.InitAdminJwtParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// 注册 API 路由
	api := r.Group("/api")
	RegisterUserRoute(api, apiController, authMiddleware)
	api.Use(authMiddleware.MiddlewareFunc()) // 验证 JWT 中间件
	RegisterUserAuthRoute(api, apiController)

	// 注册 Admin 路由
	adminApi := r.Group("/admin")
	RegisterAdminRoute(adminApi, adminController, adminAuthMiddleware)
	adminApi.Use(adminAuthMiddleware.MiddlewareFunc())
	RegisterAdminAuthRoute(adminApi, adminController)
}

// repository 初始化
func initRepository() {
	userRepository = repository.NewUserRepository(database.DB)
	adminRepository = repository.NewAdminRepository(database.DB)
}

// service 初始化
func initService() {
	userService = &service.UserService{
		UserRepository: userRepository,
	}
	adminService = &service.AdminService{
		AdminRepository: adminRepository,
	}
}

func initController() {
	apiController = &ApiControllers{
		UserController: controller.NewUserController(userService),
	}
	adminController = &AdminControllers{
		UserController:  controller.NewAdminUserController(adminService, userService),
		AdminController: controller.NewAdminController(adminService),
	}
}
