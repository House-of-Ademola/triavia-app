package main

import (
	"github.ademola.com/trivia-app/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app fiber.App) {
	app.Get("/", handlers.Home)
}
