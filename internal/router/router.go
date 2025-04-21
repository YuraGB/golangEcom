package router

import (
	"go.uber.org/zap"
	"golang-server/ent"
	"golang-server/internal/config"
	"golang-server/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, db *ent.Client, logger *zap.Logger) {
	// Middleware
	app.Use(config.RecoverConfig())
	app.Use(config.LoggerConfig())
	app.Use(config.CorsConfig())
	app.Use(config.HelmetConfig())
	app.Use(config.LimiterConfig())
	app.Use(config.CsrfConfig())

	// db connection middleware
	// This middleware will be executed for every request
	app.Use(config.DbContext(db, logger))

	// API
	api := app.Group("/api")
	RegisterUserRoutes(api)

	// Static
	app.Static("/", "./public", config.StaticConfig())

	// error handling
	app.Use(handler.ErrorHandler)
}
