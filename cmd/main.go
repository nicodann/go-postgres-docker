package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicodann/go-postgres-docker/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Blurp!")
	})

	app.Listen(":3000")
}
