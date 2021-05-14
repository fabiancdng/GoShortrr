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

type UserToCreate struct {
	Username string `json:"username" form:"username"`
	Password string `json:"-" form:"password"`
	Role     int    `json:"role" form:"role"`
}
