package handler

import (
	"golang-server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategoriesHandler (c *fiber.Ctx) error{
	list, err := service.GetAllCategories()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch category list"})
	}

	return c.JSON(list)
}