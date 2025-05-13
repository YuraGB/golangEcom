package db

import (
	"context"
	"golang-server/ent"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*ent.Client, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env не знайдено, використовую системні змінні")
	}

	DB_CONNECTION := os.Getenv("DB_CONNECTION")

	log.Println("DB_CONNECTION:", DB_CONNECTION)
	client, err := ent.Open("sqlite3", DB_CONNECTION)

	if err != nil {
		log.Fatalf("❌ Не вдалося підключитись до БД: %v", err)
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("❌ Не вдалося створити схему: %v", err)
		return nil, err
	}

	log.Println("✅ Підключення до БД успішне")
	return client, nil
}
