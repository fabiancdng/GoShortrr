package database

import (
	"database/sql"
	"log"

	"github.com/fabiancdng/GoShortrr/internal/shortlink"
	"github.com/gofiber/fiber/v2"
)

func DBConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "GoShortrr:5lVX97KMM9SbM6UH@tcp(127.0.0.1:3306)/goshortrr?parseTime=true")

	if err != nil {
		panic("Can't establish MySQL connection.")
	}

	return db
}

func Init() {
	db := DBConnection()

	_, err := db.Query("CREATE TABLE IF NOT EXISTS `goshortrr`.`shortlinks` ( `id` BIGINT NOT NULL AUTO_INCREMENT , `link` TEXT NOT NULL , `short` VARCHAR(30) NOT NULL , `user` BIGINT NOT NULL , `password` VARCHAR(50) NOT NULL , `created` TIMESTAMP NOT NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB;")

	if err != nil {
		panic(err)
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS `goshortrr`.`users` ( `user_id` INT NOT NULL AUTO_INCREMENT , `username` VARCHAR(50) NOT NULL , `password` VARCHAR(200) NOT NULL , `role` TINYINT NOT NULL , PRIMARY KEY (`user_id`)) ENGINE = InnoDB;")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func CreateShortlink(link string, short string, user int, password string) bool {
	db := DBConnection()

	_, err := db.Exec("INSERT INTO `shortlinks` (`id`, `link`, `short`, `user`, `password`, `created`) VALUES (NULL, ?, ?, ?, ?, CURRENT_TIMESTAMP());", link, short, user, password)

	defer db.Close()

	if err != nil {
		log.Println("[CREATE]", err)
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

func GetShortlink(short string) (shortlink.Shortlink, error) {
	var shortlink shortlink.Shortlink
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
