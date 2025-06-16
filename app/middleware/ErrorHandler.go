package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/pkg/response"
)

func ErrorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()
		// 请求后处理
		if len(c.Errors) > 0 {
			// TODO 是否错误记录日志
			response.Fail(c, http.StatusBadRequest, c.Errors[0].Error())
		}
	}

}
