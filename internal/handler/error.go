package handler

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx) error {
	if err := c.SendFile("./public/index.html"); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Not Found",
		})
	}
	return nil
}
