package database

import (
	"database/sql"
	"log"
)

func DBConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "GoShortrr:5lVX97KMM9SbM6UH@tcp(127.0.0.1:3306)/goshortrr")

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

	defer db.Close()
}

func CreateShortlink(link string, short string, user int, password string) bool {
	db := DBConnection()

	_, err := db.Exec("INSERT INTO `shortlinks` (`id`, `link`, `short`, `user`, `password`, `created`) VALUES (NULL, ?, ?, ?, ?, current_timestamp());", link, short, user, password)

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
	rows, err := db.Query("SELECT * FROM `shortlinks` WHERE `short` = ?", short)

	if err != nil {
		log.Println(err)
	}

	if rows.Next() {
		// Shortlink is already taken
		return false
	}

	// Shortlink is valid
	return true
}
