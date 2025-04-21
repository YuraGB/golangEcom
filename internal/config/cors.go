package config

import (
	"golang-server/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func CorsConfig() fiber.Handler {
	// Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, continue with system env")
	}

	return cors.New(cors.Config{
		AllowOrigins:     utils.GetEnv("ALLOWEDORIGINS", "*"),
		AllowMethods:     utils.GetEnv("ALLOWEDMETHODS", "GET,POST,PUT,DELETE,OPTIONS"),
		AllowHeaders:     utils.GetEnv("ALLOWEDHEADERS", "Origin,Content-Type,Accept,X-Csrf-Token,X-Requested-With"),
		AllowCredentials: utils.GetEnv("ALLOWEDCREDENTIALS", "true") == "true",
		MaxAge:           utils.AtoiDefault(utils.GetEnv("MAXAGE", "3600")),
	})
}
