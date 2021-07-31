package middlewares

import (
	"strings"

	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// A middleware that checks for an authorization header and extracts the token in it
// The token is then validated and if valid, the authorization middleware gets skipped
type TokenMiddleware struct {
	db     database.Database
	config *config.Config
	store  *session.Store
}

func (middleware *TokenMiddleware) Register(db database.Database, config *config.Config, store *session.Store, router fiber.Router) {
	middleware.db = db
	middleware.config = config
	middleware.store = store

	router.Use(middleware.execute)
}

func (middleware *TokenMiddleware) execute(ctx *fiber.Ctx) error {
	authorizationToken := ctx.Get("Authorization")
	apiAccessToken := middleware.config.WebServer.APIAccessToken

	// Skip middleware execution because there is no authorization header
	if authorizationToken == "" {
		ctx.Locals("authorized", false)
		return ctx.Next()
	}

	// Skip middleware execution because token auth isn't set up
	if apiAccessToken == "" || apiAccessToken == " " {
		return fiber.NewError(401, "Authorization Using Token Is Not Allowed")
	}

	authorizationToken = strings.ReplaceAll(authorizationToken, "Basic ", "")

	// Token is not correct
	if authorizationToken != middleware.config.WebServer.APIAccessToken {
		ctx.Locals("authorized", false)
		return ctx.Next()
	}

	ctx.Locals("authorized", true)
	// Set up a 'dummy user' with admin permissions
	// In a later version this will be replaced with the user
	// that created the token
	user := new(models.User)
	user.Id = 1
	user.Role = 1
	ctx.Locals("user", user)

	return ctx.Next()
}
