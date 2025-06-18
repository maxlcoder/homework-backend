package route

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/controller"
	"github.com/maxlcoder/homework-backend/app/middleware"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/database"
	"github.com/maxlcoder/homework-backend/pkg/response"
	"github.com/maxlcoder/homework-backend/repository"
	"github.com/maxlcoder/homework-backend/service"
)

// service 列表
var (
	userRepository repository.UserRepository
)

// service 列表
var (
	userService *service.UserService
)

var (
	apiController *ApiController
)

func ApiRoutes(r *gin.Engine) {

	initRepository()
	initService()
	initController()

	// 注册 API 路由
	api := r.Group("/api")
	// 异常处理中间件
	api.Use(middleware.ErrorHandler())

	// jwt 中间件
	authMiddleware, err := jwt.New(initJwtParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	RegisterUserRoute(api, apiController, authMiddleware)

	api.Use(handlerMiddleware(authMiddleware))
	//api.Use(authMiddleware.MiddlewareFunc())
	RegisterUserAuthRoute(api, apiController)

	// 注册 Admin 路由
	adminApi := r.Group("/admin")
	AdminRegisterUserRoute(adminApi)
}

// repository 初始化
func initRepository() {
	userRepository = repository.NewUserRepository(database.DB)
}

// service 初始化
func initService() {
	userService = &service.UserService{
		UserRepository: userRepository,
	}
}

func initController() {
	apiController = &ApiController{
		UserController: controller.NewUserController(userService),
	}
}

type lgoin struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var (
	identityKey = "id"
)

func handlerMiddleware(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func initJwtParams() *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:       "Homework",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc(),

		IdentityHandler: identityHandler(),
		Authenticator:   authenticator(),
		Authorizator:    authorizator(),
		Unauthorized:    unauthorized(),
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,

		LoginResponse: loginResponse(),
	}
}

// 负载函数
func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*model.User); ok {
			return jwt.MapClaims{
				identityKey: v.ID, // 取用户表主键作为唯一标志
			}
		}
		return jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &model.User{
			ID: claims[identityKey].(uint),
		}
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals lgoin
		if err := c.ShouldBindJSON(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := loginVals.Name
		password := loginVals.Password

		userRepository = repository.NewUserRepository(database.DB)
		userFileter := model.UserFilter{
			Name: &loginVals.Name,
		}
		user, err := userRepository.FindBy(userFileter)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}
		if model.CheckPasswordHash(password, user.Password) {
			return &model.User{
				ID:   user.ID,
				Name: userID,
			}, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*model.User); ok && v.ID == 1 {
			return true
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		response.Error(c, code, message)
	}
}

func loginResponse() func(c *gin.Context, code int, message string, time time.Time) {
	return func(c *gin.Context, code int, message string, time time.Time) {
		response.Success(c, gin.H{
			"expired": time.Format("2006-01-02 15:04:05"),
			"tokne":   message,
		})
	}
}
