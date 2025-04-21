package router

import (
	handler "golang-server/internal/handler/auth"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/login", handler.LoginHandler)
	auth.Post("/register", handler.RegisterHandler)
}
