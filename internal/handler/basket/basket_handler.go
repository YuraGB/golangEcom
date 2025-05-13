package handler

import (
	"golang-server/ent"
	"golang-server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetBasket(c *fiber.Ctx) error {
	basket := service.GetBasket()
	return c.JSON(basket)
}

func AddToBasket(c *fiber.Ctx) error {
	var item *ent.Product
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	service.AddProductToBasket(item)
	return c.JSON(service.GetBasket())
}

func RemoveFromBasket(c *fiber.Ctx) error {
	var item *ent.Product
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	service.RemoveProductFromBasket(uint8(item.ID))
	return c.JSON(service.GetBasket())
}

func ClearBasket(c *fiber.Ctx) error {
	service.ClearBasket()
	return c.JSON(fiber.Map{"message": "Basket cleared"})
}

func UpdateProductQuantity(c *fiber.Ctx) error {
	var item *ent.Product
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	// todo - add quantity field to the Product
	// service.UpdateProductQuantity(uint8(item.ID), item)
	return c.JSON(service.GetBasket())
}
