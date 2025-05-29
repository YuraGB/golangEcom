package handler

import (
	"golang-server/ent"
	"golang-server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RemoveFromBasket(c *fiber.Ctx) error {
	var item *ent.Product
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	removedItemId := service.RemoveProductFromBasket(uint8(item.ID))
	if removedItemId == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in removing item from basket"})
	}
	return c.JSON(fiber.Map{"removed_item_id": removedItemId})
}

func ClearBasket(c *fiber.Ctx) error {
	service.ClearBasket()
	return c.JSON(fiber.Map{"message": "Basket cleared"})
}
