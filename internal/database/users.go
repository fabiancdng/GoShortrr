package database

import (
	"log"

	"github.com/fabiancdng/GoShortrr/internal/models"
)

func CreateUser(userToCreate *models.UserToCreate) bool {
	db := DBConnection()

	_, err := db.Exec("INSERT INTO `users` (`user_id`, `username`, `password`, `role`, `created`) VALUES (NULL, ?, ?, ?, CURRENT_TIMESTAMP());", userToCreate.Username, userToCreate.Password, userToCreate.Role)

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

func ValidateUser(userToCreate *models.UserToCreate) int {
	if len(userToCreate.Username) < 5 {
		// Username too short
		return 801
	}

	if len(userToCreate.Username) > 49 {
		// Username too long
		return 802
	}

	if len(userToCreate.Password) < 5 {
		// Password too short
		return 803
	}

	if len(userToCreate.Password) > 199 {
		// Password too long
		return 804
	}

	// Check if username is already taken
	db := DBConnection()
	result, err := db.Query("SELECT * FROM `users` WHERE `username` = ?", userToCreate.Username)

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
