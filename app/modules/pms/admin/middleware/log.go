package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("[WMS Admin] Applying logger middleware")
		// 实际项目中可以添加日志记录逻辑
		c.Next()
	}
}
