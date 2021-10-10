package mysql

import (
	"log"

	"github.com/fabiancdng/GoShortrr/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

// Inserts the passed user into the database.
func (m *MySQL) CreateUser(user *models.User) bool {
	_, err := m.db.Exec("INSERT INTO `users` (`user_id`, `username`, `password`, `role`, `created`) VALUES (NULL, ?, ?, ?, CURRENT_TIMESTAMP());", user.Username, user.Password, user.Role)

	if err != nil {
		log.Println("[CREATE USER]", err)
		return false
	}

	return true
}

// Custom status codes for user creation
// 801 	 	Username too short
// 802 		Username too long
// 803 		Password too short
// 804 		Password too long
// 805 		Username already taken

// Validates whether or not a user is okay to be created.
func (m *MySQL) ValidateUser(user *models.User) int {
	if len(user.Username) < 5 {
		// Username too short
		return 801
	}

	if len(user.Username) > 49 {
		// Username too long
		return 802
	}

	if len(user.Password) < 5 {
		// Password too short
		return 803
	}

	if len(user.Password) > 199 {
		// Password too long
		return 804
	}

	// Check if username is already taken
	result, err := m.db.Query("SELECT * FROM `users` WHERE `username` = ?", user.Username)

	if err != nil {
		log.Println(err)
	}

	if result.Next() {
		// Username is already taken
		return 805
	}

	// User valid
	return 200
}

// Obtains a full user from the database by their username.
func (m *MySQL) GetUser(username string) (*models.User, error) {
	user := new(models.User)

	result, err := m.db.Query("SELECT * FROM `users` WHERE `username` = ?", username)
	if err != nil {
		log.Println(err)
	}

	if result.Next() {
		result.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Created,
		)

		return user, nil
	}

	return user, fiber.NewError(401, "Invalid User")
}
