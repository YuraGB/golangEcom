package main

import (
	"context"
	"golang-server/ent"
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

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("❌ Не вдалося підключитись до БД: %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("❌ Не вдалося створити схему: %v", err)
	}

	app := fiber.New()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("❌ Не вдалося ініціалізувати логер: %v", err)
	}
	defer logger.Sync()

	router.RegisterRoutes(app, client, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("🚀 Сервер запущено на порту %s", port)
	log.Fatal(app.Listen(":" + port))
}
