package main

import (
	"github.com/fabiancdng/GoShortrr/internal/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Make sure all tables exist
	database.Init()

	app := fiber.New()

	// Serve production build of React app
	app.Static("/", "./web/build")

	//////////////// ~ API ~ ////////////////

	// Route for testing purposes
	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Routes for managing shortlinks
	app.Get("/api/create")

	// Routes for managing users

	// Routes for managing keys

	app.Listen(":4000")
}
