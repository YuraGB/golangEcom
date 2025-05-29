package router

import (
	"golang-server/internal/handler/products"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductsRoutes(router fiber.Router) {
	products := router.Group("/products")
	products.Get("/", handler.GetProducts)
	products.Get("/:id", handler.GetProductById)
}
