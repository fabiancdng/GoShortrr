package webserver

import (
	"time"

	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
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

	// Session storage
	var storage = mysql.New(mysql.Config{
		Host:     config.MySQL.Host,
		Port:     config.MySQL.Port,
		Database: config.MySQL.DB,
		Username: config.MySQL.User,
		Password: config.MySQL.Password,
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

	// Routes for managing authentication / users
	apiAuthGroup.Post("/login", ws.loginUser)       // Route for checking user info and starting a session
	apiAuthGroup.Post("/register", ws.registerUser) // Route for creating a user
	apiAuthGroup.Post("/user", ws.getUser)          // Route for retrieving user info on frontend

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
