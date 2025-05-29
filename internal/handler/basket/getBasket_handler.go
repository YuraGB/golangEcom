package handler

import (
	"golang-server/internal/service"
	"golang-server/utils"
	utils_db "golang-server/utils/db"

	"github.com/gofiber/fiber/v2"
)

func GetBasket(c *fiber.Ctx) error {
	userId, err := utils.GetUserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := utils_db.GetDbFromContext(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	basket, err := service.GetBasket(userId, db, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in getting basket"})
	}
	return c.JSON(basket)
}
