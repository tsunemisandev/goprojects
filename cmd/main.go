package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsunemisandev/goprojects/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! Thiago4")
	})

	app.Listen(":3000")
}
