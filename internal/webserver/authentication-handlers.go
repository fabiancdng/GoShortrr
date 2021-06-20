package webserver

import (
	"log"

	"github.com/alexedwards/argon2id"
	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (ws *WebServer) registerUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	ctx.BodyParser(user)

	statusValid := ws.db.ValidateUser(user)

	if statusValid != 200 {
		// User is not valid
		return ctx.SendStatus(statusValid)
	}

	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)

	if err != nil {
		return ctx.SendStatus(500)
	}

	user.Password = hash

	statusCreate := ws.db.CreateUser(user)

	if statusCreate == true {
		// User has successfully been created
		return ctx.SendStatus(200)
	} else {
		// Error
		return ctx.SendStatus(500)
	}
}

func (ws *WebServer) loginUser(ctx *fiber.Ctx) error {
	login := new(models.Login)
	ctx.BodyParser(login)

	user, err := ws.db.AuthUser(*login)

	if err != nil {
		return err
	}

	sess, err := ws.store.Get(ctx)

	if err != nil {
		log.Println(err)
		return fiber.NewError(500)
	}

	sess.Set("username", user.Username)

	defer sess.Save()

	return ctx.SendStatus(200)
}

func (ws *WebServer) getUser(ctx *fiber.Ctx) error {
	user, err := ws.getUserBySession(ctx, false)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"username": user.Username,
		"role":     user.Role,
	})
}
