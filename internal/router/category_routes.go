package router

import (
	handler "golang-server/internal/handler/category"

	"github.com/gofiber/fiber/v2"
)

func RegisterCategoryRoutes(router fiber.Router) {
	categoryRouteGroup := router.Group("/category")

	categoryRouteGroup.Get("/", handler.GetCategoryByNameHandler)
	categoryRouteGroup.Get("/list", handler.GetAllCategoriesHandler)
}