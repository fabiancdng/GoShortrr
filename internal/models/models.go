package models

import "time"

// Holds data for a shortlink
type Shortlink struct {
	Id       int       `json:"id"`
	Link     string    `json:"link"`
	Short    string    `json:"short"`
	User     int       `json:"user"`
	Password bool      `json:"password"`
	Created  time.Time `json:"created"`
}

// Defines what data a shortlink create request needs to have
type ShortlinkToCreate struct {
	Link     string `json:"link"`
	Short    string `json:"short"`
	Password string `json:"password"`
}

// Holds data for a GoShortrr user
type User struct {
	Id       int       `json:"-"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     int       `json:"role"`
	Created  time.Time `json:"-"`
}

// Defines what data a login request needs to have
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
