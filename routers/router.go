package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/time-book/middleware/jwt"
	"github.com/panda8z/time-book/pkg/settings"
	"github.com/panda8z/time-book/routers/api"
	v1 "github.com/panda8z/time-book/routers/api/v1"
)

// InitR is the router`s initialization func
func InitR() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	gin.SetMode(settings.RunMode)

	// test for index router
	app.GET("/", func(ctx *gin.Context) {
		var str string
		page := ctx.Query("page")
		if page != "" {
			str = fmt.Sprintf("We recived Page number : %s", page)
		}
		ctx.JSON(http.StatusOK, gin.H{"msg": "Hello World!\n" + str})
	})

	app.Static("/static", "./web-client/static")

	// auth first
	app.GET("/auth", api.Auth)
	app.POST("/auth", api.AddAuth)
	// Api version 01 for web Ui / mobile app
	v1Api := app.Group("/api/v1")
	// use jwt middleware
	v1Api.Use(jwt.JWT())
	v1Api.GET("/note/:id", v1.Note)
	v1Api.POST("/note", v1.AddN)
	v1Api.DELETE("note/:id", v1.DeleteN)
	v1Api.PUT("/note/:id", v1.EditN)

	v1Api.GET("/test/:id", func(c *gin.Context) {
		content := fmt.Sprintf("We recived a Name: %s", c.Query("name"))
		c.JSON(http.StatusOK, gin.H{
			"msg": content,
		})
	})

	// Last middleware to match anything
	// 404 handler copy from fiber example
	app.Use(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404 了！",
		})
	})

	return app
}
