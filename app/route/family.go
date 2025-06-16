package route

import "github.com/gin-gonic/gin"

func RegisterFamilyRoute(r *gin.RouterGroup) {
	api := r.Group("families")
	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	api.GET("/:id", func(c *gin.Context) {})
	api.POST("", func(c *gin.Context) {})
}
