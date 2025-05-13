package handler

import (
	"golang-server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetCategoryByNameHandler (c *fiber.Ctx) error{
	categoryName := c.Query("name")

	if 	categoryName == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No category name was provided"})
	}

	productList, err := service.GetCategoryProducts(categoryName )


	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch category list"})
	}

	return c.JSON(productList)
}