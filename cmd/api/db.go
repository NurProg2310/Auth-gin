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
		"postgres://postgres:1234@localhost:5432/authdb?sslmode=disable",
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
	var dbname, schema, sp string
	_ = DB.QueryRow(context.Background(), "select current_database(), current_schema(), current_setting('search_path')").Scan(&dbname, &schema, &sp)
	log.Println("DB:", dbname, "schema:", schema, "search_path:", sp)
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
