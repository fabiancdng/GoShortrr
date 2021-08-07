package database

import (
	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/models"
)

// Defines what functions a database middleware must provide
type Database interface {
	// Opens a database connection
	Open(config *config.Config) error
	// Makes sure all tables exist in database
	Init() error

	////////////////
	//    USERS   //
	////////////////

	// Creates a user
	CreateUser(user *models.User) bool
	// Validates whether or not a user is okay to be created
	ValidateUser(user *models.User) int
	// Obtains a user from the database by their username
	GetUser(username string) (*models.User, error)

	/////////////////////
	//    SHORTLINKS   //
	/////////////////////

	// Creates a shortlink
	CreateShortlink(shortlinkToCreate *models.ShortlinkToCreate, user *models.User) bool
	// Validates whether or not a shortlink is okay to be created
	ValidateShortlink(short string) bool
	// Obtains a shortlink from the database by it's unique part
	GetShortlink(short string) (models.Shortlink, error)
	// Gets a list of all the user's shortlinks
	GetShortlinkList(user *models.User) ([]models.Shortlink, error)
	// Revokes/deletes a shortlink from the database by it's unique part
	DeleteShortlink(short string) (int64, error)
}
