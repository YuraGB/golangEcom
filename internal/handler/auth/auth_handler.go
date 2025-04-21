package handler

import (
	"golang-server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var userDataInput service.UserLoginInput
	if err := c.BodyParser(&userDataInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	lagged := service.LoginService(userDataInput)
	// Implement your login logic here
	return c.JSON(lagged)
}

func RegisterHandler(c *fiber.Ctx) error {
	var userDataInput service.UserRegisterInput
	if err := c.BodyParser(&userDataInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	lagged := service.RegisterService(userDataInput)
	// Implement your registration logic here
	return c.JSON(lagged)
}
