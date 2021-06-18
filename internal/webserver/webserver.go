package webserver

import (
	"time"

	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
)

type WebServer struct {
	app   *fiber.App
	store *session.Store
	db    database.Middleware
}

// Create and return a WebServer object
func NewWebServer(db database.Middleware) (*WebServer, error) {
	// Initialize Fiber app and Session store
	app := fiber.New()

	// Session storage
	var storage = mysql.New(mysql.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "goshortrr",
		Username: "goshortrr",
		Password: "5lVX97KMM9SbM6UH",
		Table:    "sessions",
		Reset:    false,
	})

	// Session middleware
	var store = session.New(session.Config{
		Storage:    storage,
		Expiration: 24 * time.Hour * 30,
	})

	// Create WebServer object and inject dependencies
	ws := &WebServer{
		app:   app,
		store: store,
		db:    db,
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

	// Routes for managing authentication / users
	apiAuthGroup.Post("/api/auth/login", ws.loginUser)       // Route for checking user info and starting a session
	apiAuthGroup.Post("/api/auth/register", ws.registerUser) // Route for creating a user
	apiAuthGroup.Post("/api/auth/user", ws.getUser)          // Route for retrieving user info on frontend

	// Sub-group that holds all routes starting with /api/shortlink/*
	apiShortlinkGroup := apiGroup.Group("/shortlink")

	// Routes for managing shortlinks
	apiShortlinkGroup.Get("/api/shortlink/get/:short", ws.getShortlink)   // Route for fetching what's behind a shortlink
	apiShortlinkGroup.Post("/api/shortlink/create", ws.createShortlink)   // Route for creating a shortlink
	apiShortlinkGroup.Delete("/api/shortlink/delete", ws.deleteShortlink) // Route for deleting a shortlink
}

// Run the Fiber webserver with all initialized routes
func (ws *WebServer) RunWebServer() error {
	// Run Fiber webserver
	err := ws.app.Listen(":4000")
	if err != nil {
		return err
	}

	return nil
}
