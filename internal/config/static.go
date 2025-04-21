package config

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func StaticConfig() fiber.Static {
	return fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: time.Second,
		MaxAge:        3600,
		Next:          nil,
	}
}
