package main

import (
	"golang-server/internal/db"
	"golang-server/internal/router"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	_ "entgo.io/ent/dialect/sql"
)

func main() {
	// init logger
	logger, errorLogger := zap.NewProduction()
	if errorLogger != nil {
		logger.Fatal("❌ Не вдалося ініціалізувати логер:", zap.Error(errorLogger))
	}
	defer logger.Sync()
	// ----------------

	// init env
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(".env не знайдено, використовую системні змінні")
	}
	// ----------------

	// db client
	client, err := db.Connect(logger)
	if err != nil {
		logger.Fatal("❌ База не підключилась: %v", zap.Error(err))
	}
	defer client.Close()
	// -------------------

	// init app + router
	app := fiber.New()

	router.RegisterRoutes(app, client, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// -----------------

	// start server
	logger.Info("🚀 Сервер запущено", zap.String("port", port))

	if err := app.Listen(":" + port); err != nil {
		logger.Fatal("❌ Помилка запуску сервера", zap.Error(err))
	}
}
