package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tsunemisandev/goprojects/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	routes := app.GetRoutes(true)

	for _, route := range routes {
		fmt.Println(route.Path)
	}

	app.Listen(":3000")
}
