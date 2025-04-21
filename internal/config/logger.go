package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerConfig() fiber.Handler {
	return logger.New()
}
