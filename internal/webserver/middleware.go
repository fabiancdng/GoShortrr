package webserver

import (
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// A middleware is a function executed before the request goes to the handler
// it can do things like checking authorization or append infos to the request.
type Middleware interface {
	// Registers a middleware to the passed router and injects all needed dependencies.
	Register(db database.Database, store *session.Store, router fiber.Router)
}
