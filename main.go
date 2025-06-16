package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/middleware"
	"github.com/maxlcoder/homework-backend/app/route"
	"github.com/maxlcoder/homework-backend/pkg/database"
	"github.com/maxlcoder/homework-backend/pkg/validator"
	"github.com/spf13/viper"
)

func setupRouter() *gin.Engine {
	// 配置项加载
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置读取失败：%s \n", err))
	}
	// 监视配置变化
	viper.WatchConfig()

	// 数据连接初始化
	err = database.InitDB()
	if err != nil {
		panic(fmt.Errorf("数据库连接初始化失败：%s \n", err))
	}

	// 参数校验翻译
	validator.InitValidator()

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// 注册路由
	api := r.Group("/api")
	// 全局异常处理中间件
	api.Use(middleware.ErrorHandler())

	// jwt 验证中间件

	adminApi := api.Group("/admin")
	route.AdminRegisterUserRoute(adminApi)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	//r.GET("/user/:name", func(c *gin.Context) {
	//	user := c.Params.ByName("name")
	//	value, ok := db[user]
	//	if ok {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	//	}
	//})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	//authorized.POST("admin", func(c *gin.Context) {
	//	user := c.MustGet(gin.AuthUserKey).(string)
	//
	//	// Parse JSON
	//	var json struct {
	//		Value string `json:"value" binding:"required"`
	//	}
	//
	//	if c.Bind(&json) == nil {
	//		db[user] = json.Value
	//		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	//	}
	//})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8083")
}
