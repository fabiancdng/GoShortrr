package webserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/webserver/controllers"
	"github.com/fabiancdng/GoShortrr/internal/webserver/middlewares"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// WebServer object holding the webserver and all its dependencies.
type WebServer struct {
	app    *fiber.App
	store  *session.Store
	config *config.Config
	db     database.Database
}

// Instantiates and returns a WebServer object.
func NewWebServer(db database.Database, config *config.Config) (*WebServer, error) {
	// Initialize Go Fiber
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// Initialize sessions middleware
	var store = session.New(session.Config{
		Expiration: 24 * time.Hour * 30,
	})

	// Create the instance and injects needed dependencies
	ws := &WebServer{
		app:    app,
		store:  store,
		config: config,
		db:     db,
	}

	ws.setup()

	return ws, nil
}

// Registers all routes and their handler functions and middlewares
// to the WebServer instance it's called from.
func (ws *WebServer) setup() {
	// Router that holds all routes starting with /api/*
	apiRouter := ws.app.Group("/api")

	// Register logging middleware
	ws.app.Use(logger.New())

	// Route for testing purposes
	apiRouter.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Register TokenMiddleware to the global /api/* router
	tokenMiddleware := new(middlewares.TokenMiddleware)
	tokenMiddleware.Register(ws.db, ws.config, ws.store, apiRouter)

	// Register AuthorizationMiddleware to the global /api/* router
	authorizationMiddleware := new(middlewares.AuthorizationMiddleware)
	authorizationMiddleware.Register(ws.db, ws.config, ws.store, apiRouter)

	// Router that holds all routes starting with /api/auth/*
	apiAuthRouter := apiRouter.Group("/auth")
	new(controllers.AuthenticationController).Register(ws.db, ws.store, apiAuthRouter)

	// Router that holds all routes starting with /api/shortlink/*
	apiShortlinkRouter := apiRouter.Group("/shortlink")
	new(controllers.ShortlinkController).Register(ws.db, ws.store, apiShortlinkRouter)

	// Serve production build of React app (and all its assets)
	ws.app.Use(filesystem.New(filesystem.Config{
		Root:         http.Dir("web/build"),
		Browse:       true,
		Index:        "index.html",
		MaxAge:       3600,
		NotFoundFile: "index.html",
	}))

	// Serve server monitor from Fiber middleware
	ws.app.Get("/monitor", monitor.New())
}

// Runs the webserver
func (ws *WebServer) RunWebServer() error {
	// Get the address and port Shortinator bind to from the config
	addressAndPort := ws.config.WebServer.AddressAndPort

	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\n>> Shortinator is now running at http://%s !\n", addressAndPort)
	color.Unset()

	// Run the Fiber webserver
	err := ws.app.Listen(addressAndPort)
	if err != nil {
		return err
	}

	return nil
}
