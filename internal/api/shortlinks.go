package api

import (
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/models"
	"github.com/fabiancdng/GoShortrr/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func GetShortlink(c *fiber.Ctx) error {
	shortlink, err := database.GetShortlink(c.Params("short"))

	if err != nil {
		return err
	}

	return c.JSON(shortlink)
}

func CreateShortlink(c *fiber.Ctx) error {
	user, err := utils.GetUserBySession(c, store, false)

	if err != nil {
		return err
	}

	shortlinkToCreate := new(models.ShortlinkToCreate)
	c.BodyParser(shortlinkToCreate)

	if shortlinkToCreate.Short == "" {
		for database.ValidateShortlink(shortlinkToCreate.Short) == false {
			short, _ := utils.GenerateShort(5)
			shortlinkToCreate.Short = short
		}
	}

	if database.ValidateShortlink(shortlinkToCreate.Short) == false {
		return fiber.NewError(409, "shortlink invalid or already taken")
	}

	database.CreateShortlink(shortlinkToCreate, user)

	return c.JSON(fiber.Map{
		"short": shortlinkToCreate.Short,
		"link":  shortlinkToCreate.Link,
	})
}

func DeleteShortlink(c *fiber.Ctx) error {
	return c.SendString("Create shortlink")
}
