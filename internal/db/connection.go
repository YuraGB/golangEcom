package db

import (
    "context"
    "database/sql"
    "golang-server/ent"
    "entgo.io/ent/dialect"
    entsql "entgo.io/ent/dialect/sql"
    "go.uber.org/zap"
    "os"

    "github.com/joho/godotenv"
    _ "modernc.org/sqlite"
)

func Connect(logger *zap.Logger) (*ent.Client, error) {
    err := godotenv.Load()
    if err != nil {
        logger.Warn(".env не знайдено, використовую системні змінні")
    }

    DB_CONNECTION := os.Getenv("DB_CONNECTION")
    if DB_CONNECTION == "" {
        DB_CONNECTION = "file:ent.db"
    }
    logger.Info("DB_CONNECTION", zap.String("connection_string", DB_CONNECTION))

    db, err := sql.Open("sqlite", DB_CONNECTION)
    if err != nil {
        logger.Fatal("Не вдалося відкрити з'єднання з БД", zap.Error(err))
        return nil, err
    }

    // Вручну увімкнути foreign keys
    _, err = db.Exec("PRAGMA foreign_keys = ON;")
    if err != nil {
        logger.Fatal("Не вдалося увімкнути foreign_keys pragma", zap.Error(err))
        return nil, err
    }

    drv := entsql.OpenDB(dialect.SQLite, db)
    client := ent.NewClient(ent.Driver(drv))

    if err := client.Schema.Create(context.Background()); err != nil {
        logger.Fatal("Не вдалося створити схему", zap.Error(err))
        return nil, err
    }

    logger.Info("Підключення до БД успішне")
    return client, nil
}
