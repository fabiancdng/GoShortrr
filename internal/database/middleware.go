package database

import (
	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/models"
)

// Define what functions a database middleware must provide
type Middleware interface {
	// Open a database connection
	Open(config *config.Config) error
	// Make sure all tables exist in database
	Init() error

	////////////////
	//    USERS   //
	////////////////

	// Create a user
	CreateUser(user *models.User) bool
	// Validate whether or not a user is okay to be created
	ValidateUser(user *models.User) int
	// Look up user (using credentials) and return if existing
	AuthUser(login models.Login) (models.User, error)
	// Return a user without having to provide credentials
	GetUser(username string) (models.User, error)

	/////////////////////
	//    SHORTLINKS   //
	/////////////////////

	// Create a shortlink
	CreateShortlink(shortlinkToCreate *models.ShortlinkToCreate, user *models.User) bool
	// Validate whether or not a shortlink is okay to be created
	ValidateShortlink(short string) bool
	// Get a shortlink
	GetShortlink(short string) (models.Shortlink, error)
}
