package database

import (
	"log"

	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(user *models.User) bool {
	db := DBConnection()

	_, err := db.Exec("INSERT INTO `users` (`user_id`, `username`, `password`, `role`, `created`) VALUES (NULL, ?, ?, ?, CURRENT_TIMESTAMP());", user.Username, user.Password, user.Role)

	defer db.Close()

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

func ValidateUser(user *models.User) int {
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
	db := DBConnection()
	result, err := db.Query("SELECT * FROM `users` WHERE `username` = ?", user.Username)

	if err != nil {
		log.Println(err)
	}

	if result.Next() {
		// Username is already taken
		return 805
	}

	defer db.Close()

	// User valid
	return 200
}

func GetUser(short string) (models.Shortlink, error) {
	var shortlink models.Shortlink
	var shortlinkPassword string
	db := DBConnection()

	result, err := db.Query("SELECT * FROM `shortlinks` WHERE `short` = ?", short)

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

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
