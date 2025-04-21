package handler

import (
	"github.com/gofiber/fiber/v2"
	"golang-server/internal/service"
)

func GetBasket(c *fiber.Ctx) error {
	basket := service.GetBasket()
	return c.JSON(basket)
}

func AddToBasket(c *fiber.Ctx) error {
	var item service.Product
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	service.AddProductToBasket(item)
	return c.JSON(service.GetBasket())
}

func RemoveFromBasket(c *fiber.Ctx) error {
	var item service.Product
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	service.RemoveProductFromBasket(item.ID)
	return c.JSON(service.GetBasket())
}

func ClearBasket(c *fiber.Ctx) error {
	service.ClearBasket()
	return c.JSON(fiber.Map{"message": "Basket cleared"})
}

func UpdateProductQuantity(c *fiber.Ctx) error {
	var item service.Product
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	service.UpdateProductQuantity(item.ID, item.Quantity)
	return c.JSON(service.GetBasket())
}
