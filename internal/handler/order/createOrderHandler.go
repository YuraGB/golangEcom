package handler

import (
	"context"
	"golang-server/internal/models"
	"golang-server/internal/service"
	"golang-server/utils"
	utils_db "golang-server/utils/db"

	"github.com/gofiber/fiber/v2"
)

func CreateOrderHandler(c *fiber.Ctx) error {
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
	// Parse the request body into the order variable
	var orderRequestInput models.OrderRequestInput
	if err := c.BodyParser(&orderRequestInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input" + err.Error()})
	}

	basketProducts, err := service.GetBasket(userId, db, context.Background())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "there is no products in the basket"})
	}

	total := 0.0
	// Calculate the total price from the order products
	for _, product := range basketProducts {
		total += product.Price * float64(product.Quantity)
	}

	// Check if the total price is valid
	if total <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Total price must be greater than zero"})
	}

	// Create the order input model
	var orderInput = models.OrderInput{
		Order: models.Order{
			PaymentType:   orderRequestInput.PaymentMethod,
			Address:       orderRequestInput.Address,
			City:          orderRequestInput.City,
			State:         orderRequestInput.State,
			Zip:           orderRequestInput.Zip,
			UserID:        orderRequestInput.UserId,
			OrderProducts: orderRequestInput.OrderProducts,
		},
		TotalPrice: total,
		Status:     "NEW", // Default status for new orders
	}

	// Create the order
	createdOrder, err := service.CreateOrder(userId, db, c.Context(), orderInput)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in creating order" + err.Error()})
	}
	// Return the created order
	return c.Status(fiber.StatusCreated).JSON(createdOrder)
}
