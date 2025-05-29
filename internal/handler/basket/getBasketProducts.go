package handler

import (
	"golang-server/internal/service"
	"golang-server/utils"
	utils_db "golang-server/utils/db"

	"github.com/gofiber/fiber/v2"
)

func GetBasketProducts(c *fiber.Ctx) error {
	// Get user ID from context
	userId, err := utils.GetUserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Get database connection from context
	db, err := utils_db.GetDbFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Get the user's basket
	basket, err := service.GetBasket(userId, db, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in getting basket"})
	}

	if len(basket) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Basket is empty"})
	}

	// Get the first basket item (or handle it as needed)
	products, err := service.GetBasketProducts(basket)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in getting products"})
	}

	return c.JSON(products)
}
