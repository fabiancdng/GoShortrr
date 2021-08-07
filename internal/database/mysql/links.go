package mysql

import (
	"log"

	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Creates a shortlink
func (m *MySQL) CreateShortlink(shortlinkToCreate *models.ShortlinkToCreate, user *models.User) bool {
	_, err := m.db.Exec("INSERT INTO `shortlinks` (`id`, `link`, `short`, `user`, `password`, `created`) VALUES (NULL, ?, ?, ?, ?, CURRENT_TIMESTAMP());", shortlinkToCreate.Link, shortlinkToCreate.Short, user.Id, shortlinkToCreate.Password)

	if err != nil {
		log.Println("[CREATE LINK]", err)
		return false
	}

	return true
}

// Validates whether or not a shortlink is okay to be created
func (m *MySQL) ValidateShortlink(short string) bool {
	if short == "" {
		// Shortlink can't be empty
		return false
	}

	if len(short) > 30 {
		// Shortlink too long
		return false
	}

	// Checks if shortlink is already taken
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

// Obtains a user from the database by it's unique part
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

	return shortlink, fiber.NewError(404, "Shortlink Not Found")
}

// Gets a list of all the user's shortlinks
func (m *MySQL) GetShortlinkList(user *models.User) ([]models.Shortlink, error) {
	shortlink := new(models.Shortlink)
	var shortlinkList []models.Shortlink
	var shortlinkPassword string

	result, err := m.db.Query("SELECT * FROM `shortlinks` WHERE `user` = ?", user.Id)
	if err != nil {
		return shortlinkList, fiber.NewError(500)
	}

	for result.Next() {
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

		shortlinkList = append(shortlinkList, *shortlink)
	}

	return shortlinkList, nil
}

func (m *MySQL) DeleteShortlink(short string) (int64, error) {
	result, err := m.db.Exec("DELETE FROM `shortlinks` WHERE `short`=?", short)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
