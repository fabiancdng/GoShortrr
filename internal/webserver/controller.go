package webserver

import (
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// A controller registers all routes and handlers
// for a specific part of the API
// It also holds the handler functions for that part
type Controller interface {
	// Register registers a controller to the passed fiber.Router
	// which typically is a so called 'group'
	Register(db database.Middleware, store *session.Store, router fiber.Router)
}
