package database

import (
	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/models"
)

// Defines what methods a database middleware must provide.
type Database interface {
	//////////////////////////////////////
	//                                  //
	//    INTERNAL DATABASE AFFAIRES    //
	//                                  //
	//////////////////////////////////////

	// Opens a database connection that is safe for concurrent use
	// as it utilizes a connection pool.
	//
	// Reference: https://pkg.go.dev/database/sql#Open
	Open(config *config.Config) error

	// Checks whether all needed tables exist and if not,
	// it automatically creates them as well as an admin user.
	Init() error

	//////////////////////////////////
	//                              //
	//       USERS / ACCOUNTS       //
	//                              //
	//////////////////////////////////

	// Inserts the passed user into the database.
	CreateUser(user *models.User) bool

	// Validates whether or not a user is okay to be created.
	ValidateUser(user *models.User) int

	// Obtains a full user from the database by their username.
	GetUser(username string) (*models.User, error)

	////////////////////////////
	//                        //
	//       SHORTLINKS       //
	//                        //
	////////////////////////////

	// Inserts the passed shortlink into the database and
	// therefore finalizes its creation.
	CreateShortlink(shortlinkToCreate *models.ShortlinkToCreate, user *models.User) bool

	// Performs a DB lookup for the passed unique part of a shortlink
	// and checks whether or not it is already taken by a shortlink
	IsShortlinkTaken(short string) bool

	// Obtains a shortlink from the database by it's unique part.
	GetShortlink(short string) (models.Shortlink, error)

	// Returns a list of all the user's shortlinks.
	GetShortlinkList(user *models.User) ([]models.Shortlink, error)

	// Revokes/deletes a shortlink from the database.
	// The shortlink is identified by its unique part.
	DeleteShortlink(short string) (int64, error)
}
