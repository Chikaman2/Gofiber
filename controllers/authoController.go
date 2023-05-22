package controllers

import (
	"github.com/chikaman2/Gofiber/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	user := models.Users{
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.JSON(user)
}
