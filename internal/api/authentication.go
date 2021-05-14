package api

import (
	"log"
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
	Host:     "127.0.0.1",
	Port:     3306,
	Database: "goshortrr",
	Username: "goshortrr",
	Password: "5lVX97KMM9SbM6UH",
	Table:    "sessions",
	Reset:    false,
})

// Sessions
var store = session.New(session.Config{
	Storage:    storage,
	Expiration: 24 * time.Hour * 30,
})

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	c.BodyParser(user)

	statusValid := database.ValidateUser(user)

	if statusValid != 200 {
		// User is not valid
		return c.SendStatus(statusValid)
	}

	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)

	if err != nil {
		return c.SendStatus(500)
	}

	user.Password = hash

	statusCreate := database.CreateUser(user)

	if statusCreate == true {
		// User has successfully been created
		return c.SendStatus(200)
	} else {
		// Error
		return c.SendStatus(500)
	}
}

func LoginUser(c *fiber.Ctx) error {
	login := new(models.Login)
	c.BodyParser(login)

	user, err := database.GetUser(*login)

	if err != nil {
		return err
	}

	sess, err := store.Get(c)

	if err != nil {
		log.Println(err)
		return fiber.NewError(500)
	}

	sess.Set("username", user.Username)

	defer sess.Save()

	return c.SendStatus(200)
}
