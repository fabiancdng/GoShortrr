package middlewares

import (
	"log"

	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// A middleware that checks for a valid session and passes the authorization
// status as well as the user to the next middleware/handler.
type AuthorizationMiddleware struct {
	db     database.Database
	config *config.Config
	store  *session.Store
}

// Registers this middleware to the passed router and injects all needed dependencies.
func (middleware *AuthorizationMiddleware) Register(db database.Database, config *config.Config, store *session.Store, router fiber.Router) {
	middleware.db = db
	middleware.config = config
	middleware.store = store

	// Register middleware function the the router
	router.Use(middleware.execute)
}

// Run this middleware and pass on result(s).
func (middleware *AuthorizationMiddleware) execute(ctx *fiber.Ctx) error {
	// Skip middleware because request is already authorized
	if ctx.Locals("authorized") == true {
		return ctx.Next()
	}

	sess, err := middleware.store.Get(ctx)
	if err != nil {
		log.Println(err)
		return fiber.NewError(500)
	}

	username := sess.Get("username")
	if username == nil {
		ctx.Locals("authorized", false)
		return ctx.Next()
	}

	user, err := middleware.db.GetUser(username.(string))
	if err != nil {
		ctx.Locals("authorized", false)
		return ctx.Next()
	}

	// Pass authorization status and user to next middleware/handler
	ctx.Locals("authorized", true)
	ctx.Locals("user", user)

	return ctx.Next()
}
