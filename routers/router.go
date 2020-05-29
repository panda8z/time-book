package routers

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func InitRouter() *fiber.App {
	app := fiber.New()

	app.Get("/" , func(ctx *fiber.Ctx) {
		var str string
		page := ctx.Query("page")
		if page != "" {
			str = fmt.Sprintf("We recived Page number : %s", page)
		}
		ctx.SendString("Hello World!\n" + str)
	})
	v1Api := app.Group("/api/v1")
	v1Api.Get("/test", func(c *fiber.Ctx) {
		content := fmt.Sprintf("We recived a Name: %s", c.Query("name"))
		c.SendString(content)
	})

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
		c.SendString("404 了！")
	})

	return app
}
