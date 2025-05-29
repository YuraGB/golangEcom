package handler

import (
	"golang-server/internal/models"
	"golang-server/internal/service"
	"golang-server/utils"
	utils_db "golang-server/utils/db"

	"github.com/gofiber/fiber/v2"
)

func UpdateProductQuantity(c *fiber.Ctx) error {
	var item *models.BasketProduct
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	userId, err := utils.GetUserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := utils_db.GetDbFromContext(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedItem := service.UpdateProductQuantity(userId, db, c.Context(), []models.BasketProduct{*item})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in getting basket"})
	}

	return c.JSON(updatedItem)
}
