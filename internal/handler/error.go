package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx) error {
	log.Println("error")
	if err := c.SendFile("./public/index.html"); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Not Found",
		})
	}
	return nil
}
