package request

import (
	"fmt"
	"strings"

	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/maxlcoder/homework-backend/pkg/validator"
)

type PageRequest struct {
	Page    int `form:"page" binding:"omitempty,min=1" label:"页码" default:"1"`
	PerPage int `form:"per_page" binding:"omitempty,min=1,max=100" label:"每页数量" default:"100"` // 最大支持 100 TODO 配置项
}

type IdRequest struct {
	Id uint `json:"id" binding:"required,gt=0"`
}

func BindAndSetDefaults(c *gin.Context, req interface{}) error {
	// 绑定校验之前给默认值
	// 判断一下请求方式
	ct := c.ContentType()
	if ct == "application/json" {
		if c.Request.Method == "GET" {
			if err := c.ShouldBindWith(req, binding.Query); err != nil {
				errorTrans := validator.TranslateError(err)
				return fmt.Errorf("%s", strings.Join(errorTrans, ","))

			}
		} else if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			if err := c.ShouldBind(req); err != nil {
				errorTrans := validator.TranslateError(err)
				return fmt.Errorf("%s", strings.Join(errorTrans, ","))

			}
		}
	} else {
		if err := c.ShouldBind(req); err != nil {
			errorTrans := validator.TranslateError(err)
			return fmt.Errorf("%s", strings.Join(errorTrans, ","))
		}
	}
	// 应用默认值（用 creasty/defaults 或你自己写的 applyDefaults）
	if err := defaults.Set(req); err != nil {
		return err
	}
	return nil
}
