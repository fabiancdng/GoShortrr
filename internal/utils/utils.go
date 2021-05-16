package utils

import (
	"crypto/rand"
	"log"

	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateShort(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

func GetUserBySession(c *fiber.Ctx, store *session.Store, admin bool) (*models.User, error) {
	sessionCookie := c.Cookies("session_id")
	user := new(models.User)

	// Check if request is authorized
	if sessionCookie == "" {
		return user, fiber.NewError(401, "no session")
	} else {
		// Check if user has sufficient permissions
		sess, err := store.Get(c)

		if err != nil {
			log.Println(err)
			return user, fiber.NewError(500)
		}

		username := sess.Get("username")

		if username == nil {
			return user, fiber.NewError(401, "invalid session")
		}

		userDB, err := database.GetUser(username.(string))
		user = &userDB

		if err != nil {
			return user, fiber.NewError(401, "invalid session")
		}

		if admin == true {
			if user.Role < 1 {
				return user, fiber.NewError(401, "insufficient permissions")
			}
		}
	}

	return user, nil
}
