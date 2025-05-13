package utils_db

import (
	"golang-server/ent"
	"golang-server/internal/config"

	"github.com/gofiber/fiber/v2"
)

func GetDbFromContext(c *fiber.Ctx) (*ent.Client, error) {
	ctx, ok := c.Locals("ctx").(*config.AppContext)
	if !ok {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "AppContext not found")
	}

	db := ctx.DB
	if db == nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "DB not found in context")
	}

	return db, nil
}