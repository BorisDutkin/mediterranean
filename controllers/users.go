package controllers

import (
	"github.com/BorisDutkin/mediterranean/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {		
	u := models.User {
		FirstName: "Boris",
		LastName: "Dutkin",
		Email: "boris.d@gmail.com",
		Password: "123456789",
	}

	return c.Status(fiber.StatusOK).JSON(u)
}