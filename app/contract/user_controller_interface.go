package contract

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(c *gin.Context)
}

type AdminUserController interface {
	List(c *gin.Context)
}
