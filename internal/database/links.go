package database

import (
	"log"

	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/gofiber/fiber/v2"
)

func Init() {
	db := DBConnection()

	_, err := db.Query("CREATE TABLE IF NOT EXISTS `goshortrr`.`shortlinks` ( `id` BIGINT NOT NULL AUTO_INCREMENT , `link` TEXT NOT NULL , `short` VARCHAR(30) NOT NULL , `user` BIGINT NOT NULL , `password` VARCHAR(50) NOT NULL , `created` TIMESTAMP NOT NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB;")

	if err != nil {
		panic(err)
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS `goshortrr`.`shortlinks` ( `id` BIGINT NOT NULL AUTO_INCREMENT , `link` TEXT NOT NULL , `short` VARCHAR(30) NOT NULL , `user` INT NOT NULL , `password` VARCHAR(50) NOT NULL , `created` TIMESTAMP NOT NULL , PRIMARY KEY (`id`), FOREIGN KEY (`user`) REFERENCES `users`(`user_id`)) ENGINE = InnoDB;")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func CreateShortlink(shortlinkToCreate *models.ShortlinkToCreate, user *models.User) bool {
	db := DBConnection()

	_, err := db.Exec("INSERT INTO `shortlinks` (`id`, `link`, `short`, `user`, `password`, `created`) VALUES (NULL, ?, ?, ?, ?, CURRENT_TIMESTAMP());", shortlinkToCreate.Link, shortlinkToCreate.Short, user.Id, shortlinkToCreate.Password)

	defer db.Close()

	if err != nil {
		log.Println("[CREATE LINK]", err)
		return false
	}

	return true
}

func ValidateShortlink(short string) bool {
	if short == "" {
		// Shortlink can't be empty
		return false
	}

	if len(short) > 30 {
		// Shortlink too long
		return false
	}

	// Check if shortlink is already taken
	db := DBConnection()
	result, err := db.Query("SELECT * FROM `shortlinks` WHERE `short` = ?", short)

	if err != nil {
		log.Println(err)
	}

	if result.Next() {
		// Shortlink is already taken
		return false
	}

	defer db.Close()

	// Shortlink is valid
	return true
}

func GetShortlink(short string) (models.Shortlink, error) {
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
