package webserver

import (
	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/fabiancdng/GoShortrr/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func (ws *WebServer) getShortlink(ctx *fiber.Ctx) error {
	shortlink, err := ws.db.GetShortlink(ctx.Params("short"))

	if err != nil {
		return err
	}

	return ctx.JSON(shortlink)
}

func (ws *WebServer) createShortlink(ctx *fiber.Ctx) error {
	user, err := utils.GetUserBySession(ctx, ws.db, ws.store, false)

	if err != nil {
		return err
	}

	shortlinkToCreate := new(models.ShortlinkToCreate)
	ctx.BodyParser(shortlinkToCreate)

	if shortlinkToCreate.Short == "" {
		for ws.db.ValidateShortlink(shortlinkToCreate.Short) == false {
			short, _ := utils.GenerateShort(5)
			shortlinkToCreate.Short = short
		}
	}

	if ws.db.ValidateShortlink(shortlinkToCreate.Short) == false {
		return fiber.NewError(409, "shortlink invalid or already taken")
	}

	ws.db.CreateShortlink(shortlinkToCreate, user)

	return ctx.JSON(fiber.Map{
		"short": shortlinkToCreate.Short,
		"link":  shortlinkToCreate.Link,
	})
}

func (ws *WebServer) deleteShortlink(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete shortlink")
}
