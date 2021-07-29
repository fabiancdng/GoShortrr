package middlewares

import (
	"log"

	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// A middleware that checks for a BasicAuth token or a valid session
// and passes the authorization status as well as the user to the next middleware/handler
type AuthorizationMiddleware struct {
	db    database.Database
	store *session.Store
}

func (middleware *AuthorizationMiddleware) Register(db database.Database, store *session.Store, router fiber.Router) {
	middleware.db = db
	middleware.store = store

	// Registers middleware function the the router
	router.Use(middleware.execute)
}

func (middleware *AuthorizationMiddleware) execute(ctx *fiber.Ctx) error {
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
