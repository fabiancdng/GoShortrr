package mysql

import (
	"log"

	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Inserts the passed shortlink into the database and
// therefore finalizes its creation.
func (m *MySQL) CreateShortlink(shortlinkToCreate *models.ShortlinkToCreate, user *models.User) bool {
	_, err := m.db.Exec("INSERT INTO `shortlinks` (`id`, `link`, `short`, `user`, `created`) VALUES (NULL, ?, ?, ?, CURRENT_TIMESTAMP());", shortlinkToCreate.Link, shortlinkToCreate.Short, user.Id)

	if err != nil {
		log.Println("[CREATE LINK]", err)
		return false
	}

	return true
}

// Performs a DB lookup for the passed unique part of a shortlink
// and checks whether or not it is already taken by a shortlink
func (m *MySQL) IsShortlinkTaken(short string) bool {
	// Checks whether or not shortlink is already taken
	result, err := m.db.Query("SELECT * FROM `shortlinks` WHERE `short` = ?", short)
	if err != nil {
		log.Println(err)
	}

	if result.Next() {
		// Shortlink is already taken
		return true
	}

	// Shortlink is not taken
	// Hence, it's okay to be created
	return false
}

// Obtains a shortlink from the database by it's unique part.
func (m *MySQL) GetShortlink(short string) (models.Shortlink, error) {
	var shortlink models.Shortlink

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
			&shortlink.Created,
		)

		return shortlink, nil
	}

	return shortlink, fiber.NewError(404, "Shortlink Not Found")
}

// Returns a list of all the user's shortlinks.
func (m *MySQL) GetShortlinkList(user *models.User) ([]models.Shortlink, error) {
	shortlink := new(models.Shortlink)
	var shortlinkList []models.Shortlink

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
			&shortlink.Created,
		)

		shortlinkList = append(shortlinkList, *shortlink)
	}

	return shortlinkList, nil
}

// Revokes/deletes a shortlink from the database.
// The shortlink is identified by its unique part.
func (m *MySQL) DeleteShortlink(short string) (int64, error) {
	result, err := m.db.Exec("DELETE FROM `shortlinks` WHERE `short`=?", short)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
