package webserver

import (
	"time"

	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/webserver/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type WebServer struct {
	app    *fiber.App
	store  *session.Store
	config *config.Config
	db     database.Middleware
}

// Create and return a WebServer object
func NewWebServer(db database.Middleware, config *config.Config) (*WebServer, error) {
	// Initialize Fiber app and Session store
	app := fiber.New()

	// Session middleware
	var store = session.New(session.Config{
		Expiration: 24 * time.Hour * 30,
	})

	// Create WebServer object and inject dependencies
	ws := &WebServer{
		app:    app,
		store:  store,
		config: config,
		db:     db,
	}

	ws.registerHandlers()

	// Return created WebServer object
	return ws, nil
}

// Register all routes and their handler functions
func (ws *WebServer) registerHandlers() {
	// +++++++++ STATIC ++++++++++++++++++

	// Serve production build of React app
	ws.app.Static("/*", "./web/build")

	// Serve server monitor from Fiber middleware
	ws.app.Get("/monitor", monitor.New())

	// +++++++++ API ++++++++++++++++++

	// Group that holds all routes starting with /api/*
	apiGroup := ws.app.Group("/api")

	// Route for testing purposes
	apiGroup.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Sub-group that holds all routes starting with /api/auth/*
	apiAuthGroup := apiGroup.Group("/auth")
	new(controllers.AuthenticationController).Register(ws.db, ws.store, apiAuthGroup)

	// Sub-group that holds all routes starting with /api/shortlink/*
	apiShortlinkGroup := apiGroup.Group("/shortlink")

	// Routes for managing shortlinks
	apiShortlinkGroup.Get("/get/:short", ws.getShortlink)   // Route for looking up what's behind a shortlink
	apiShortlinkGroup.Post("/create", ws.createShortlink)   // Route for creating a shortlink
	apiShortlinkGroup.Delete("/delete", ws.deleteShortlink) // Route for deleting a shortlink
}

// Run the Fiber webserver with all initialized routes
func (ws *WebServer) RunWebServer() error {
	// Run Fiber webserver
	err := ws.app.Listen(ws.config.WebServer.AddressAndPort)
	if err != nil {
		return err
	}

	return nil
}
