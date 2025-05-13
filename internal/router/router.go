package router

import (
	"golang-server/ent"
	"golang-server/internal/config"
	"golang-server/internal/handler"
	"log"

	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, db *ent.Client, logger *zap.Logger) {
	
	// Middleware
	app.Use(config.RecoverConfig())
	app.Use(config.LoggerConfig())
	app.Use(config.CorsConfig())
	app.Use(config.HelmetConfig())
	app.Use(config.LimiterConfig())
	// app.Use(config.CsrfConfig())

	// db connection middleware
	app.Use(config.DbContext(db, logger))

	// API
	api := app.Group("/api")
	RegisterUserRoutes(api)
	RegisterBasketRoutes(api)
	RegisterAuthRoutes(api)
	RegisterProductsRoutes(api)
	RegisterPostRoutes(api)
	RegisterCategoryRoutes(api)

	// non-case "/api"
	app.Use("/api", func(c *fiber.Ctx) error {
		log.Print(c.Request().URI())
		return c.Status(404).JSON(fiber.Map{
			"error": "API route not found",
		})
	})

	// Static
	app.Static("/", "./public", config.StaticConfig())

	// error handling
	app.Use(handler.ErrorHandler)

	// non-case
	app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})

}
