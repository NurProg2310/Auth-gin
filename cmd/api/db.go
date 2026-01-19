package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// InitDB инициализирует подключение к PostgreSQL
func InitDB() {
	dsn := getEnv(
		"DATABASE_URL",
		"postgres://auth:1234@localhost:5432/authdb?sslmode=disable",
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("DB connect error: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}

	DB = pool
	log.Println("✅ DB connected")
}

// CloseDB корректно закрывает пул
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// getEnv читает переменную окружения или возвращает default
func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
