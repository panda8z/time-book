package main

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/:name/:age?", func(c *fiber.Ctx) {
		str := fmt.Sprintf("First route: [name]: %s [age]: %s\n", c.Params("name"), c.Params("age"))
		c.SendString(str)
		c.Next() // Continue the stack
	})

	app.Get("/api/*", func(c *fiber.Ctx) {
		str := fmt.Sprintf("Last route:  [*]: %s\n", c.Params("*"))
		c.SendString(str)
	})

	app.Get("/panda/*", func(c *fiber.Ctx) {
		str := fmt.Sprintf("Lastest route:  [*]: %s\n", c.Params("*"))
		c.SendString(str)
	})

	app.Listen(3000)
	// GET /api/register
	// First route: [name]: api [age]: register
	// Last route:  [*]: register
}
