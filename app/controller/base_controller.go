package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/pkg/response"
)

type BaseController struct{}

func (controller *BaseController) Success(c *gin.Context, data interface{}) {
	response.Success(c, data)
}

func (controller *BaseController) Forbidden(c *gin.Context, msg string) {
	response.Forbidden(c, msg)
}

func (controller *BaseController) Error(c *gin.Context, code int, msg string) {
	response.Error(c, code, msg)
}

// 参数校验
