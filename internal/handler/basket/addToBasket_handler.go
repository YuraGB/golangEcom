package handler

import (
	"golang-server/internal/models"
	"golang-server/internal/service"
	"golang-server/utils"
	"log"

	utils_db "golang-server/utils/db"

	"github.com/gofiber/fiber/v2"
)

func AddToBasket(c *fiber.Ctx) error {
	var item *models.BasketProduct

	// Parse the request body into the item variable
	c.BodyParser(&item)
	log.Println("Item", item)

	if item == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Get the user ID from the context
	userId, err := utils.GetUserIdFromCtx(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid access token"})
	}

	// Get the database connection from the context
	db, err := utils_db.GetDbFromContext(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database connection error")
	}

	basket, err := service.GetBasket(userId, db, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in getting basket"})
	}
	if basket == nil {
		// If the user doesn't have a basket, create a new one
		// and add the item to it
		items := []models.BasketProduct{*item}
		newBasket, err := service.CreateBasket(userId, db, c.Context(), items)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in creating new basket"})
		}
		return c.JSON(newBasket)
	}

	addedItem, err := service.AddProductToBasket(userId, db, c.Context(), *item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in adding product to basket"})
	}
	// Return the updated basket
	return c.JSON(addedItem)
}
