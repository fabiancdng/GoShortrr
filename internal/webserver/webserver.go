package webserver

import (
	"log"
	"time"

	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/webserver/controllers"
	"github.com/fabiancdng/GoShortrr/internal/webserver/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type WebServer struct {
	app    *fiber.App
	store  *session.Store
	config *config.Config
	db     database.Database
}

// Creates, sets up and returns a WebServer
func NewWebServer(db database.Database, config *config.Config) (*WebServer, error) {
	// Initialize Fiber app and Session store
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// Session middleware
	var store = session.New(session.Config{
		Expiration: 24 * time.Hour * 30,
	})

	// Create WebServer and injects dependencies
	ws := &WebServer{
		app:    app,
		store:  store,
		config: config,
		db:     db,
	}

	ws.setup()

	// Return the created WebServer object
	return ws, nil
}

// Registers all routes and their handler functions
func (ws *WebServer) setup() {
	////////////////////
	//     STATIC     //
	////////////////////

	// Serve production build of React app
	ws.app.Static("/*", "./web/build")
	// Serve server monitor from Fiber middleware
	ws.app.Get("/monitor", monitor.New())

	// Register logging middleware
	ws.app.Use(logger.New())

	/////////////////
	//     API     //
	/////////////////

	// Router that holds all routes starting with /api/*
	apiRouter := ws.app.Group("/api")

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

}

// Runs the webserver
func (ws *WebServer) RunWebServer() error {
	log.Println(">> Webserver is now running!")
	// Run the Fiber webserver
	err := ws.app.Listen(ws.config.WebServer.AddressAndPort)
	if err != nil {
		return err
	}

	return nil
}
