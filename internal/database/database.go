package database

import (
	"database/sql"
	"log"
)

func DBConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "GoShortrr:5lVX97KMM9SbM6UH@tcp(127.0.0.1:3306)/goshortrr")

	if err != nil {
		log.Panic("Couldn't establish MySQL connection.")
	}

	return db
}

func Init() {
	db := DBConnection()

	_, err := db.Query("CREATE TABLE IF NOT EXISTS `goshortrr`.`shortlinks` ( `id` BIGINT NOT NULL AUTO_INCREMENT , `link` TEXT NOT NULL , `short` VARCHAR(30) NOT NULL , `user` BIGINT NOT NULL , `password` VARCHAR(50) NOT NULL , `created` TIMESTAMP NOT NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB;")

	if err != nil {
		panic(err)
	}

}
