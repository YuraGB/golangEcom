package main

import (
	"golang-server/internal/db"
	"golang-server/internal/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env не знайдено, використовую системні змінні")
	}

	client, nil := db.Connect()

	if err != nil {
		log.Fatalf("❌ База не підключилась: %v", err)
	}

	defer client.Close()

	app := fiber.New()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("❌ Не вдалося ініціалізувати логер: %v", err)
	}
	defer logger.Sync()

	router.RegisterRoutes(app, client, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Сервер запущено на порту %s", port)
	log.Fatal(app.Listen(":" + port))
}
