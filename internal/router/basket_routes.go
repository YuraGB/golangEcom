package router

import (
	handler "golang-server/internal/handler/basket"

	"github.com/gofiber/fiber/v2"
)

func RegisterBasketRoutes(router fiber.Router) {
	basket := router.Group("/basket")
	basket.Get("/", handler.GetBasket)
	basket.Post("/add", handler.AddToBasket)
	basket.Delete("/remove", handler.RemoveFromBasket)
	basket.Delete("/clear", handler.ClearBasket)
	basket.Patch("/update", handler.UpdateProductQuantity)
}
