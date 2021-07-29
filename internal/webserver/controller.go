package webserver

import (
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// A Controller holds all routes and handlers
// for a specific part of the API (for instance /api/auth/*)
type Controller interface {
	// Register registers a controller's routes and handlers to the passed fiber.Router
	Register(db database.Database, store *session.Store, router fiber.Router)
}
