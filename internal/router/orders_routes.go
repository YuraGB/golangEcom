package router

import (
	handler "golang-server/internal/handler/order"

	"github.com/gofiber/fiber/v2"
)

func RegisterOrderRoutes(router fiber.Router) {
	orderRouteGroup := router.Group("/order")

	orderRouteGroup.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Order route is working",
		})

	})
	// orderRouteGroup.Get("/:id", GetOrderByIdHandler)
	orderRouteGroup.Post("/create/", handler.CreateOrderHandler)
	orderRouteGroup.Post("/create", handler.CreateOrderHandler)

}
