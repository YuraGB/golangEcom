package handler

import (
	"golang-server/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	products, err := service.GetAllProducts()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}

	return c.JSON(products)
}

func GetProductById (c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("asdasd")
	product, err := service.GetProduct(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch product"})
	}

	return c.JSON(product)
}
