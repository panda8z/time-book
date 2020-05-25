package main

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	// Match any route
	app.Use(func(c *fiber.Ctx) {
		fmt.Println("First middleware")
		c.Next()
	})

	// Match all routes starting with /api
	app.Use("/api", func(c *fiber.Ctx) {
		fmt.Println("Second middleware")
		c.Next()
	})

	// GET /api/register
	app.Get("/api/list", func(c *fiber.Ctx) {
		fmt.Println("Last middleware")
		c.Send("Hello, World!")
	})

	//
	//app.Get("/:name/:age?", func(c *fiber.Ctx) {
	//	str := fmt.Sprintf("First route: [name]: %s [age]: %s\n", c.Params("name"), c.Params("age"))
	//	c.SendString(str)
	//	c.Next() // Continue the stack
	//})
	//
	//app.Get("/api/*", func(c *fiber.Ctx) {
	//	str := fmt.Sprintf("Last route:  [*]: %s\n", c.Params("*"))
	//	c.SendString(str)
	//})
	//
	//app.Get("/panda/*", func(c *fiber.Ctx) {
	//	str := fmt.Sprintf("Lastest route:  [*]: %s\n", c.Params("*"))
	//	c.SendString(str)
	//})
	//
	//// Last middleware to match anything
	//app.Use(func(c *fiber.Ctx) {
	//	c.SendStatus(404) // => 404 "Not Found"
	//	c.SendString("404 not found anthing")
	//})

	app.Listen(3000)
	// GET /api/register
	// First route: [name]: api [age]: register
	// Last route:  [*]: register
}
