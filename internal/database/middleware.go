package database

import (
	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/models"
)

// Defines what functions a database middleware must provide
type Middleware interface {
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
	// Looks up user (using credentials) and return if existing
	AuthUser(login models.Login) (*models.User, error)
	// Returns a user without having to provide credentials
	GetUser(username string) (*models.User, error)

	/////////////////////
	//    SHORTLINKS   //
	/////////////////////

	// Creates a shortlink
	CreateShortlink(shortlinkToCreate *models.ShortlinkToCreate, user *models.User) bool
	// Validates whether or not a shortlink is okay to be created
	ValidateShortlink(short string) bool
	// Gets a shortlink
	GetShortlink(short string) (models.Shortlink, error)
}
