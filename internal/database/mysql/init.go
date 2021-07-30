package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/fabiancdng/GoShortrr/internal/config"
)

// MySQL database middleware
// This implements the database.Database interface
type MySQL struct {
	db     *sql.DB
	config *config.Config
}

// Makes sure all tables exist in database
func (m *MySQL) Init() error {
	// Create the users table if it doesn't exist
	_, err := m.db.Exec("CREATE TABLE IF NOT EXISTS `users` ( `user_id` INT NOT NULL AUTO_INCREMENT , `username` VARCHAR(50) NOT NULL , `password` VARCHAR(200) NOT NULL , `role` TINYINT NOT NULL , `created` TIMESTAMP NOT NULL , PRIMARY KEY (`user_id`)) ENGINE = InnoDB;")
	if err != nil {
		panic(err)
	}

	// Create the shortlinks table if it doesn't exist
	_, err = m.db.Exec("CREATE TABLE IF NOT EXISTS `shortlinks` ( `id` BIGINT NOT NULL AUTO_INCREMENT , `link` TEXT NOT NULL , `short` VARCHAR(30) NOT NULL , `user` INT NOT NULL , `password` VARCHAR(50) NOT NULL , `created` TIMESTAMP NOT NULL , PRIMARY KEY (`id`), FOREIGN KEY (`user`) REFERENCES `users`(`user_id`)) ENGINE = InnoDB;")
	if err != nil {
		panic(err)
	}

	log.Println(">> MySQL tables have been initialized!")

	var userCount int
	var rows *sql.Rows
	rows, err = m.db.Query("SELECT COUNT(*) FROM `users`;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&userCount)
	}

	if userCount == 0 {
		log.Println(">> There is no registered user!\n--> Creating a temporary admin account with username 'admin' and password 'admin'")
		_, err := m.db.Exec("INSERT INTO `users` (`user_id`, `username`, `password`, `role`, `created`) VALUES (NULL, 'admin', '$argon2id$v=19$m=16,t=2,p=1$MjNtZjg3MmtmOQ$1OGCDVpPbrxhjEV8YRh0Kw', 1, CURRENT_TIMESTAMP());")
		if err != nil {
			panic(err)
		}
	} else {
		log.Printf(">> Currently, there are %s registered users", strconv.Itoa(userCount))
	}

	return nil
}

// Opens a database connection
func (m *MySQL) Open(config *config.Config) error {
	m.config = config

	dbConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		m.config.MySQL.User,
		m.config.MySQL.Password,
		m.config.MySQL.Host,
		strconv.Itoa(m.config.MySQL.Port),
		m.config.MySQL.DB,
	)

	var err error
	m.db, err = sql.Open("mysql", dbConfig)

	if err != nil {
		panic("Can't establish MySQL connection.")
	}

	log.Println(">> Successfully established connection to MySQL server!")
	return nil
}
