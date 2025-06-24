package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 请求前
		c.Next()
	}
}
