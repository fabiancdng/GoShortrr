package controllers

import (
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/fabiancdng/GoShortrr/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// The controller for handling all requests to /api/shortlink/*
// These routes are for managing shortlinks
type ShortlinkController struct {
	db    database.Database
	store *session.Store
}

func (controller *ShortlinkController) Register(db database.Database, store *session.Store, router fiber.Router) {
	controller.db = db
	controller.store = store

	// Route for looking up what's behind a shortlink
	router.Get("/get/:short", controller.getShortlink)
	// Route for creating a shortlink
	router.Post("/create", controller.createShortlink)
	// Route for deleting a shortlink
	router.Delete("/delete/:short", controller.deleteShortlink)
}

// Returns data behind a shortlink
func (controller *ShortlinkController) getShortlink(ctx *fiber.Ctx) error {
	shortlink, err := controller.db.GetShortlink(ctx.Params("short"))

	if err != nil {
		return err
	}

	return ctx.JSON(shortlink)
}

// Creates a shortlink
func (controller *ShortlinkController) createShortlink(ctx *fiber.Ctx) error {
	if ctx.Locals("authorized") == false {
		return fiber.NewError(401)
	}

	// Get user from the request's locals
	user := ctx.Locals("user").(*models.User)

	shortlinkToCreate := new(models.ShortlinkToCreate)
	ctx.BodyParser(shortlinkToCreate)

	if shortlinkToCreate.Short == "" {
		for controller.db.ValidateShortlink(shortlinkToCreate.Short) == false {
			short, _ := utils.GenerateShort(5)
			shortlinkToCreate.Short = short
		}
	}

	if controller.db.ValidateShortlink(shortlinkToCreate.Short) == false {
		return fiber.NewError(409, "Shortlink Invalid Or Already Taken")
	}

	controller.db.CreateShortlink(shortlinkToCreate, user)

	return ctx.JSON(fiber.Map{
		"short": shortlinkToCreate.Short,
		"link":  shortlinkToCreate.Link,
	})
}

// Deletes a shortlink
func (controller *ShortlinkController) deleteShortlink(ctx *fiber.Ctx) error {
	if ctx.Locals("authorized") == false {
		return fiber.NewError(401)
	}

	// Get user from the request's locals
	user := ctx.Locals("user").(*models.User)
	short := ctx.Params("short")

	shortlink, err := controller.db.GetShortlink(short)
	if err != nil {
		return fiber.NewError(404)
	}

	// If shortlink was created by a different user
	if shortlink.User != user.Id {
		// Check if user has admin permissions to delete the shortlink anyway
		if user.Role < 1 {
			return fiber.NewError(403, "Insufficient Permissions")
		}
	}

	var affected int64
	affected, err = controller.db.DeleteShortlink(shortlink.Short)
	if err != nil {
		return fiber.NewError(500)
	}

	if affected < 1 {
		return fiber.NewError(500)
	}

	return ctx.SendStatus(200)
}
