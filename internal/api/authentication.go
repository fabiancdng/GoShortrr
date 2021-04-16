package api

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
)

// Storage for Fiber (mainly for sessions)
var storage = mysql.New(mysql.Config{
	Host:       "127.0.0.1",
	Port:       3306,
	Database:   "goshortrr",
	Username:   "goshortrr",
	Password:   "5lVX97KMM9SbM6UH",
	Table:      "fiber_storage",
	Reset:      false,
	GCInterval: 10 * time.Second,
})

// Sessions
var store = session.New(session.Config{
	Storage: storage,
})

func CreateUser(c *fiber.Ctx) error {
	log.Println("Test")
	return nil
}
