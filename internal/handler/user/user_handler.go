package handler

import (
	"github.com/gofiber/fiber/v2"
	"golang-server/internal/config"
	"golang-server/internal/service"
)

func GetUsers(c *fiber.Ctx) error {
	ctx := c.Locals("ctx").(*config.AppContext)
	users := service.GetAllUsers(ctx.DB)
	return c.JSON(users)
}
