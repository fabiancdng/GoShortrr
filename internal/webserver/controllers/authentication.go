package controllers

import (
	"log"

	"github.com/alexedwards/argon2id"
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// The controller for handling all requests to /api/auth/*
// These routes are for managing authentication/users
type AuthenticationController struct {
	db    database.Database
	store *session.Store
}

func (controller *AuthenticationController) Register(db database.Database, store *session.Store, router fiber.Router) {
	controller.db = db
	controller.store = store

	// Route for checking user info and starting a session
	router.Post("/login", controller.loginUser)
	// Route for creating a user
	router.Post("/register", controller.registerUser)
	// Route for retrieving user info on frontend
	router.Post("/user", controller.getUser)
}

func (controller *AuthenticationController) registerUser(ctx *fiber.Ctx) error {
	if ctx.Locals("authorized") == false {
		return fiber.NewError(401)
	}

	// Get user from the request's locals
	user := ctx.Locals("user").(*models.User)

	// Return an error if user doesn't have admin permissions
	if user.Role != 1 {
		return fiber.NewError(401, "insufficient permissions")
	}

	userToRegister := new(models.User)
	ctx.BodyParser(userToRegister)

	statusValid := controller.db.ValidateUser(userToRegister)

	if statusValid != 200 {
		// User is not valid
		return ctx.SendStatus(statusValid)
	}

	hash, err := argon2id.CreateHash(userToRegister.Password, argon2id.DefaultParams)

	if err != nil {
		return ctx.SendStatus(500)
	}

	userToRegister.Password = hash

	statusCreate := controller.db.CreateUser(userToRegister)

	if statusCreate == true {
		// User has successfully been created
		return ctx.SendStatus(200)
	} else {
		// Error
		return ctx.SendStatus(500)
	}
}

func (controller *AuthenticationController) loginUser(ctx *fiber.Ctx) error {
	login := new(models.Login)
	ctx.BodyParser(login)

	user, err := controller.db.GetUser(login.Username)
	if err != nil {
		return fiber.NewError(401)
	}

	match, err := argon2id.ComparePasswordAndHash(login.Password, user.Password)
	if err != nil {
		return fiber.NewError(401)
	}

	if match != true {
		return fiber.NewError(401)
	}

	sess, err := controller.store.Get(ctx)
	if err != nil {
		log.Println(err)
		return fiber.NewError(500)
	}

	sess.Set("username", user.Username)

	defer sess.Save()

	return ctx.SendStatus(200)
}

func (controller *AuthenticationController) getUser(ctx *fiber.Ctx) error {
	if ctx.Locals("authorized") == false {
		return fiber.NewError(401)
	}

	// Gets user from the request's locals
	user := ctx.Locals("user").(*models.User)

	return ctx.JSON(fiber.Map{
		"username": user.Username,
		"role":     user.Role,
	})
}
