package handler

import (
	"golang-server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetSearchResults(c *fiber.Ctx) error {
	searchParams := c.Query("q")

	if searchParams == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "There is no search query provided"})
	}

	searchResults, error := service.GetSearchProducts(searchParams)

	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in search"})

	}

	return c.JSON(searchResults)
}
