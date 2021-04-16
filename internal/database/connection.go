package database

import "database/sql"

func DBConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "GoShortrr:5lVX97KMM9SbM6UH@tcp(127.0.0.1:3306)/goshortrr?parseTime=true")

	if err != nil {
		panic("Can't establish MySQL connection.")
	}

	return db
}
