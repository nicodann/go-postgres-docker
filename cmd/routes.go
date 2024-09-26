package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicodann/go-postgres-docker/handlers"
)

func setupRoutes(app *fiber.App) {
	// app.Get("/", handlers.Home)
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
