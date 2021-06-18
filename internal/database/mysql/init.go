package mysql

import "database/sql"

// Create MySQL object
// This implements the database.Middleware interface
type MySQL struct {
	db *sql.DB
}

// Make sure all tables exist in database
func (m *MySQL) Init() error {
	// Create the users table if it doesn't exist
	_, err := m.db.Query("CREATE TABLE IF NOT EXISTS `goshortrr`.`users` ( `user_id` INT NOT NULL AUTO_INCREMENT , `username` VARCHAR(50) NOT NULL , `password` VARCHAR(200) NOT NULL , `role` TINYINT NOT NULL , `created` TIMESTAMP NOT NULL , PRIMARY KEY (`user_id`)) ENGINE = InnoDB;")
	if err != nil {
		panic(err)
	}

	// Create the shortlinks table if it doesn't exist
	_, err = m.db.Query("CREATE TABLE IF NOT EXISTS `goshortrr`.`shortlinks` ( `id` BIGINT NOT NULL AUTO_INCREMENT , `link` TEXT NOT NULL , `short` VARCHAR(30) NOT NULL , `user` INT NOT NULL , `password` VARCHAR(50) NOT NULL , `created` TIMESTAMP NOT NULL , PRIMARY KEY (`id`), FOREIGN KEY (`user`) REFERENCES `users`(`user_id`)) ENGINE = InnoDB;")
	if err != nil {
		panic(err)
	}

	return nil
}

// Open a database connection
func (m *MySQL) Open() error {
	var err error
	m.db, err = sql.Open("mysql", "GoShortrr:5lVX97KMM9SbM6UH@tcp(127.0.0.1:3306)/goshortrr?parseTime=true")

	if err != nil {
		panic("Can't establish MySQL connection.")
	}

	return nil
}
