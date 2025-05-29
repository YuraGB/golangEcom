package handler

import (
	"golang-server/internal/models"
	"golang-server/internal/service"
	"golang-server/utils"
	utils_db "golang-server/utils/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func MergeBaskets(c *fiber.Ctx) error {
	var items *[]models.BasketProduct

	if err := c.BodyParser(&items); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	log.Println("Items", items)
	userId, err := utils.GetUserIdFromCtx(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid access token")
	}

	db, err := utils_db.GetDbFromContext(c)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database connection error")
	}

	// Check if the user has an existing basket
	// If not, create a new one
	existingBasket, err := service.GetBasket(userId, db, c.Context())
	log.Println("Existing Basket", existingBasket)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in getting existing basket"})
	}

	if existingBasket == nil {
		newBasket, err := service.CreateBasket(userId, db, c.Context(), *items)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in creating new basket"})
		}

		return c.JSON(newBasket)
	}

	// If the user has an existing basket, merge the new items into it
	mergedBasket, err := service.MergeBaskets(userId, *items, existingBasket, db, c.Context())
	log.Println("Merged Basket", mergedBasket)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in merging baskets"})
	}

	updatedBasket, err := service.GetBasket(userId, db, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in getting updated basket"})
	}

	// Return the merged basket
	return c.JSON(updatedBasket)
}
