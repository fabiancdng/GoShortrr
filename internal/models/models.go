package models

import "time"

type Shortlink struct {
	Id       int       `json:"id"`
	Link     string    `json:"link"`
	Short    string    `json:"short"`
	User     string    `json:"user"`
	Password bool      `json:"password"`
	Created  time.Time `json:"created"`
}

type ShortlinkToCreate struct {
	ApiKey   string `json:"-" form:"key"`
	Link     string `json:"link" form:"link"`
	Short    string `json:"short" form:"short"`
	Password string `json:"-" form:"password"`
}

type User struct {
	Id       int       `json:"-"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     int       `json:"role"`
	Created  time.Time `json:"-"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
