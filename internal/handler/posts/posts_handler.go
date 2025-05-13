package handler

import (
	"golang-server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	posts, err := service.GetPosts()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch prosts"})
	}

	return c.JSON(posts)
}