package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type WebServer struct {
	app   *fiber.App
	store *session.Store
}

func NewWebServer() (*WebServer, error) {
	app := fiber.New()
	store := session.New()

	ws := &WebServer{
		app:   app,
		store: store,
	}

	return ws, nil
}
