package middleware

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/pkg/response"
)

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 角色信息获取
		adminRoleName, ok := c.Get("admin_role_name")
		fmt.Println(adminRoleName)
		if !ok {
			response.Unauthorized(c, "请重新登录")
			c.Abort()
			return
		}
		method := c.Request.Method
		path := c.FullPath()
		fmt.Println(path)

		// 获取当前用户的角色、
		adminRoleNameStr, ok := adminRoleName.(string)
		if !ok {
			adminRoleNameStr = ""
		}
		ok, err := e.Enforce("role_"+adminRoleNameStr, "1", path, method)
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
