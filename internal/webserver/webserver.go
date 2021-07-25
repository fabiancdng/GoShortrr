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

// Creates, sets up and returns a WebServer
func NewWebServer(db database.Middleware, config *config.Config) (*WebServer, error) {
	// Initializes Fiber app and Session store
	app := fiber.New()

	// Session middleware
	var store = session.New(session.Config{
		Expiration: 24 * time.Hour * 30,
	})

	// Creates WebServer and injects dependencies
	ws := &WebServer{
		app:    app,
		store:  store,
		config: config,
		db:     db,
	}

	ws.setup()

	// Returns created WebServer object
	return ws, nil
}

// Registers all routes and their handler functions
func (ws *WebServer) setup() {
	////////////////////
	//     STATIC     //
	////////////////////

	// Serves production build of React app
	ws.app.Static("/*", "./web/build")
	// Serves server monitor from Fiber middleware
	ws.app.Get("/monitor", monitor.New())

	/////////////////
	//     API     //
	/////////////////

	// Router that holds all routes starting with /api/*
	apiGroup := ws.app.Group("/api")

	// Route for testing purposes
	apiGroup.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Router that holds all routes starting with /api/auth/*
	apiAuthRouter := apiGroup.Group("/auth")
	new(controllers.AuthenticationController).Register(ws.db, ws.store, apiAuthRouter)

	// Router that holds all routes starting with /api/shortlink/*
	apiShortlinkRouter := apiGroup.Group("/shortlink")
	new(controllers.ShortlinkController).Register(ws.db, ws.store, apiShortlinkRouter)

}

// Runs the webserver
func (ws *WebServer) RunWebServer() error {
	// Runs the Fiber webserver
	err := ws.app.Listen(ws.config.WebServer.AddressAndPort)
	if err != nil {
		return err
	}

	return nil
}
