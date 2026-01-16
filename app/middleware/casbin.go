package middleware

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/pkg/response"
)

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 角色信息获取
		value, ok := c.Get("login_admin_role")
		if !ok {
			response.Unauthorized(c, "请重新登录")
			c.Abort()
			return
		}
		method := c.Request.Method
		path := c.FullPath()
		fmt.Println(path)

		// 获取当前用户的角色、
		role, ok := value.(model.Role)
		if !ok {
			response.Unauthorized(c, "权限不足")
			c.Abort()
			return
		}
		ok, err := e.Enforce(fmt.Sprintf("role_%d", role.ID), fmt.Sprintf("%d", role.TenantId), path, method)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		if !ok {
			response.Error(c, http.StatusForbidden, "权限不足")
			c.Abort()
			return
		}
		c.Next()
	}
}
