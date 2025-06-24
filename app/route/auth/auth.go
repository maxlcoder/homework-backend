package auth

import (
	"log"
	"net/http"
	"reflect"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/database"
	"github.com/maxlcoder/homework-backend/pkg/response"
	"github.com/maxlcoder/homework-backend/repository"
)

var (
	userRepository repository.UserRepository
)

type lgoin struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var (
	identityKey = "id"
)

func HandlerMiddleware(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func InitJwtParams() *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:       "Homework",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc[model.User, *model.User](),

		IdentityHandler: identityHandler(),
		Authenticator:   authenticator[model.User, *model.User](),
		Authorizator:    authorizator[model.User](),
		Unauthorized:    unauthorized(),
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,

		LoginResponse: loginResponse(),
	}
}

func InitAdminJwtParams() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "Homework",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc[model.Admin, *model.Admin](),

		IdentityHandler: identityHandler(),
		Authenticator:   authenticator[model.Admin, *model.Admin](),
		Authorizator:    authorizator[model.Admin](),
		Unauthorized:    unauthorized(),
		TokenLookup:     "header: Authorization", // 请求 token 设置，支持多种 header: Authorization, query: token, cookie: jwt
		TimeFunc:        time.Now,

		LoginResponse: loginResponse(),
	}
}

// 负载函数
func payloadFunc[T any, PT interface {
	*T
	model.Authenticatable
}]() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(PT); ok {
			userType := reflect.TypeOf(new(T)).Elem().Name()
			return jwt.MapClaims{
				identityKey: v.GetId(), // 取用户表主键作为唯一标志
				"user_type": userType,
			}
		}
		return jwt.MapClaims{}
	}
}

func authenticator[T any, PT interface {
	*T
	model.Authenticatable
}]() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals lgoin
		if err := c.ShouldBindJSON(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := loginVals.Name
		password := loginVals.Password
		log.Println("login", userID)
		tType := reflect.TypeOf(new(T)).Elem().Name()
		log.Println("tType", tType)
		query := database.DB.Model(new(T)).Where("name = ?", userID)
		user, err := repository.First[T, PT](query)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		if model.CheckPasswordHash(password, user.GetPassword()) {
			return user, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}

func authorizator[T model.User | model.Admin]() func(data interface{}, c *gin.Context) bool {
	// 用户类型传入
	return func(data interface{}, c *gin.Context) bool {
		tType := reflect.TypeOf(new(T)).Elem().Name()
		dataType := reflect.TypeOf(data).Elem().Name()
		if tType != dataType {
			return false
		}
		if _, ok := data.(*T); !ok {
			return false
		}
		// 获取 data 值
		return true
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		response.Error(c, code, "未授权："+message)
	}
}

func loginResponse() func(c *gin.Context, code int, message string, time time.Time) {
	return func(c *gin.Context, code int, message string, time time.Time) {
		response.Success(c, gin.H{
			"expired": time.Format("2006-01-02 15:04:05"),
			"token":   message,
		})
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		// 类型转换
		userIdFloat, ok := claims[identityKey].(float64)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "无效的用户 ID")
			return nil
		}
		userType := claims["user_type"]
		userId := uint(userIdFloat)
		switch userType {
		case "User":
			return &model.User{
				ID: userId,
			}
		case "Admin":
			return &model.Admin{
				ID: userId,
			}
		}
		return nil
	}
}
