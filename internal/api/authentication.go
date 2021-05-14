package api

import (
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/models"
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
	userToCreate := new(models.UserToCreate)
	c.BodyParser(userToCreate)

	statusValid := database.ValidateUser(userToCreate)

	if statusValid != 200 {
		// User is not valid
		return c.SendStatus(statusValid)
	}

	hash, err := argon2id.CreateHash(userToCreate.Password, argon2id.DefaultParams)

	if err != nil {
		return c.SendStatus(500)
	}

	userToCreate.Password = hash

	statusCreate := database.CreateUser(userToCreate)

	if statusCreate == true {
		// User has been created
		return c.SendStatus(200)
	} else {
		// Error
		return c.SendStatus(500)
	}
}
