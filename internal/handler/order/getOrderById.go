package handler

import (
	"golang-server/internal/service"
	"golang-server/utils"
	utils_db "golang-server/utils/db"

	"github.com/gofiber/fiber/v2"
)

func GetOrderByIdHandler(c *fiber.Ctx) error {
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

	// Get order ID from URL parameters
	orderId := c.Params("id")
	if orderId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Order ID is required"})
	}
	// Get the order by ID
	order, err := service.GetOrderById(userId, db, c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in getting order"})
	}
	if order == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
	}

	// Return the order
	return c.JSON(order)
}
