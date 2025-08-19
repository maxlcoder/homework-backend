package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/middleware"
	"github.com/maxlcoder/homework-backend/app/route"
	"github.com/maxlcoder/homework-backend/database"
	"github.com/maxlcoder/homework-backend/database/seed"
	"github.com/maxlcoder/homework-backend/pkg/validator"
	"github.com/maxlcoder/homework-backend/service"
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

	// casbin 初始化
	enforcer, err := service.NewCasbin(database.DB)
	if err != nil {
		panic(fmt.Errorf("Casbin 初始化失败: %s \n", err))
	}

	// 参数校验翻译
	validator.InitValidator()

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// 替换 Gin JSON 渲染器
	//r.JSON = jsoniter.ConfigCompatibleWithStandardLibrary

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 全局中间件
	r.Use(middleware.ErrorHandler())
	route.ApiRoutes(r, enforcer)

	err = seed.InitSeed(database.DB, r, enforcer)
	if err != nil {
		panic(fmt.Errorf("数据库初始化失败：%s \n", err))
	}

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8083")
}
