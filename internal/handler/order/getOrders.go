package handler

import (
    "golang-server/internal/service"
    "golang-server/utils"
    utils_db "golang-server/utils/db"

    "github.com/gofiber/fiber/v2"
)

func GetOrdersHandler(c *fiber.Ctx) error {
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

    // Get the user's orders
    orders, err := service.GetOrders(userId, db, c.Context())

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in getting orders"})
    }
    if len(orders) == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No orders found"})
    }
    // Return the orders
    return c.JSON(orders)
}
    