package router

import (
	"golang-server/internal/handler/search"

	"github.com/gofiber/fiber/v2"
)

func RegisterSearchRoute(router fiber.Router) {
	router.Get("/search", handler.GetSearchResults)	
}