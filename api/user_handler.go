package api

import (
	"hotel-reservation/types"

	"github.com/gofiber/fiber/v2"
)

// start func name with upper case to make it public

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "At the watercooler",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
