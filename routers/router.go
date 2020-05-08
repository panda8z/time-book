package router

import (
	"github.com/gin-gonic/gin"
	"github.com/panda8z/time-book/pkg/setting"
)

// InitRouter with no warnning
func InitRouter() *gin.Engine {
	engine := New()
	gin.SetMode(setting.RunMode)
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	return engine
}
