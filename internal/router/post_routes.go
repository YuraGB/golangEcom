package router

import (
	handler "golang-server/internal/handler/posts"

	"github.com/gofiber/fiber/v2"
)

func RegisterPostRoutes(router fiber.Router) {
	posts := router.Group("/posts")

	posts.Get("/", handler.GetPosts)
}