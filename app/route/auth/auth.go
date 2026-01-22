package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	core_model "github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/database"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/pkg/response"
	"github.com/maxlcoder/homework-backend/repository"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type lgoin struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var (
	identityKey = "id"
)

func InitMiddleware(authMiddleware *jwt.GinJWTMiddleware) {
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}

func InitJwtParams() *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:       "Homework",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc[core_model.User, *core_model.User](),

		IdentityHandler: identityHandler(),
		Authenticator:   authenticator[core_model.User, *core_model.User](),
		Authorizator:    authorizator[core_model.User](),
		Unauthorized:    unauthorized(),
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,

		LoginResponse: loginResponse(),
	}
}

func InitAdminJwtParams() *jwt.GinJWTMiddleware {
	// 获取 jwt 过期时间配置
	timeout := viper.GetDuration("jwt.timeout")
	if timeout == 0 {
		timeout = time.Hour
	}
	return &jwt.GinJWTMiddleware{
		Realm:       "Homework",
		Key:         []byte("secret key"),
		Timeout:     timeout,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc[core_model.Admin, *core_model.Admin](),

		IdentityHandler: identityHandler(),
		Authenticator:   authenticator[core_model.Admin, *core_model.Admin](),
		Authorizator:    authorizator[core_model.Admin](),
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

func authorizator[T core_model.User | core_model.Admin]() func(data interface{}, c *gin.Context) bool {
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
		fmt.Println(userType)
		userId := uint(userIdFloat)
		switch userType {
		case "User":
			// 设置全局 user_id
			c.Set("user_id", userId)
			var user core_model.User
			user.ID = userId
			return &user
		case "Admin":
			// 设置全局 admin_id
			c.Set("login_admin_id", userId)
			// 管理员角色一对多，当前角色存储在 admin 表中，先设置是否为超管简化后续判断
			ctx := context.Background()
			admin, error := gorm.G[core_model.Admin](database.DB).Where("id = ?", userId).First(ctx)
			if error != nil {
				response.Error(c, http.StatusUnauthorized, "当前用户信息异常")
				return nil
			}
			c.Set("login_admin_role_id", admin.RoleId)
			// 获取管理员角色
			if admin.RoleId > 0 {
				role, error := gorm.G[core_model.Role](database.DB).Where("id = ?", admin.RoleId).First(ctx)
				if error != nil {
					response.Error(c, http.StatusUnauthorized, "当前用户角色信息异常")
					return nil
				}
				c.Set("login_admin_role", role)
			}
			role, exists := c.Get("login_admin_role")
			if exists {
				fmt.Println("login_admin_role:", role)
			}
			admin.ID = userId
			return &admin
		}
		return nil
	}
}
