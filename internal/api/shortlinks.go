package api

import (
	"github.com/fabiancdng/GoShortrr/internal/database"
	"github.com/fabiancdng/GoShortrr/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type ShortlinkToCreate struct {
	ApiKey   string `json:"-" form:"key"`
	Link     string `json:"link" form:"link"`
	Short    string `json:"short" form:"short"`
	Password string `json:"-" form:"password"`
}

func GetShortlink(c *fiber.Ctx) error {
	return c.SendString("Get shortlink")
}

func CreateShortlink(c *fiber.Ctx) error {
	shortlinkToCreate := new(ShortlinkToCreate)
	c.BodyParser(shortlinkToCreate)

	if shortlinkToCreate.Short == "" {
		for database.ValidateShortlink(shortlinkToCreate.Short) == false {
			shortlinkToCreate.Short, _ = utils.GenerateShort(5)
		}
	}

	if database.ValidateShortlink(shortlinkToCreate.Short) == false {
		return fiber.NewError(409, "shortlink already taken")
	}

	database.InsertShortlink(
		shortlinkToCreate.Link,
		shortlinkToCreate.Short,
		1,
		shortlinkToCreate.Password,
	)

	return c.JSON(shortlinkToCreate)
}

func DeleteShortlink(c *fiber.Ctx) error {
	return c.SendString("Create shortlink")
}
