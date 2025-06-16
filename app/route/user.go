package route

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/controller"
	"github.com/maxlcoder/homework-backend/pkg/database"
	"github.com/maxlcoder/homework-backend/repository"
	"github.com/maxlcoder/homework-backend/service"
)

func RegisterUserRoute(r *gin.RouterGroup, handlerFunc gin.HandlerFunc) {

	// repository 初始化
	userRepository := repository.NewUserRepository(database.DB)

	// 服务初始化
	userService := &service.UserService{
		UserRepository: userRepository,
	}

	// 初始化控制器，注入服务
	userController := controller.NewUserController(userService)
	r.POST("users:register", userController.Register) // 注册
	r.POST("login", userController.Login)             // 登录

	auth := r.Group("")
	auth.Use(handlerFunc)
	auth.GET("me", userController.Me) // 个人中心
}

func AdminRegisterUserRoute(r *gin.RouterGroup) {
	userService := &service.UserService{}
	userController := controller.NewUserController(userService)
	r.GET("users", userController.Register)
}


func initParams() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "homework",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		IdentityKey:   "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {}
	}
}