package config

import (
	"golang-server/ent"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AppContext struct {
	DB     *ent.Client
	Logger *zap.Logger
	// etc
}

func DbContext(db *ent.Client, logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := &AppContext{
			DB:     db,
			Logger: logger,
		}
		c.Locals("ctx", ctx)
		return c.Next()
	}
}
