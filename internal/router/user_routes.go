package router

import (
	"golang-server/internal/handler/user"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/", handler.GetUsers)
}
