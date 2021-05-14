/*
	++++++++++++++++++++++++++ GoShortrr ++++++++++++++++++++++++++

	A fast, simple and powerful URL Shortener built with Go and React.

	Copyright (c) 2021 Fabian R. (fabiancdng)

	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
*/

package main

import (
	"github.com/fabiancdng/GoShortrr/internal/api"
	"github.com/fabiancdng/GoShortrr/internal/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	// Make sure all tables exist
	database.Init()

	app := fiber.New()

	// Serve production build of React app
	app.Static("/", "./web/build")

	// Serve server monitor from Fiber middleware
	app.Get("/monitor", monitor.New())

	//////////////// ~ API ~ ////////////////

	// Route for testing purposes
	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Routes for managing shortlinks
	app.Get("/api/shortlink/get/:short", api.GetShortlink)
	app.Post("/api/shortlink/create", api.CreateShortlink)
	app.Delete("/api/shortlink/delete", api.DeleteShortlink)

	// Routes for managing authentication / users
	app.Post("/api/auth/login", api.LoginUser)
	app.Post("/api/auth/register", api.RegisterUser)

	app.Listen(":4000")
}
