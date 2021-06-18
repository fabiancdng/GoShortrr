package mysql

import (
	"log"

	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Create a shortlink
func (m *MySQL) CreateShortlink(shortlinkToCreate *models.ShortlinkToCreate, user *models.User) bool {
	_, err := m.db.Exec("INSERT INTO `shortlinks` (`id`, `link`, `short`, `user`, `password`, `created`) VALUES (NULL, ?, ?, ?, ?, CURRENT_TIMESTAMP());", shortlinkToCreate.Link, shortlinkToCreate.Short, user.Id, shortlinkToCreate.Password)

	if err != nil {
		log.Println("[CREATE LINK]", err)
		return false
	}

	return true
}

// Validate whether or not a shortlink is okay to be created
func (m *MySQL) ValidateShortlink(short string) bool {
	if short == "" {
		// Shortlink can't be empty
		return false
	}

	if len(short) > 30 {
		// Shortlink too long
		return false
	}

	// Check if shortlink is already taken
	result, err := m.db.Query("SELECT * FROM `shortlinks` WHERE `short` = ?", short)

	if err != nil {
		log.Println(err)
	}

	if result.Next() {
		// Shortlink is already taken
		return false
	}

	// Shortlink is valid
	return true
}

// Get a shortlink
func (m *MySQL) GetShortlink(short string) (models.Shortlink, error) {
	var shortlink models.Shortlink
	var shortlinkPassword string

	result, err := m.db.Query("SELECT * FROM `shortlinks` WHERE `short` = ?", short)

	if err != nil {
		log.Println(err)
	}

	if result.Next() {
		result.Scan(
			&shortlink.Id,
			&shortlink.Link,
			&shortlink.Short,
			&shortlink.User,
			&shortlinkPassword,
			&shortlink.Created,
		)

		if shortlinkPassword == "" {
			shortlink.Password = false
		} else {
			shortlink.Password = true
		}

		return shortlink, nil
	}

	return shortlink, fiber.NewError(404, "shortlink not found")
}
