package api

import (
	"fmt"
	"hotel-reservation/customTypes"

	"github.com/gofiber/fiber/v2"
)

func getAuthUser(c *fiber.Ctx) (*customTypes.User, error) {
	user, ok := c.Context().UserValue("user").(*customTypes.User)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return user, nil
}
