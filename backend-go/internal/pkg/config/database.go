package config

import (
	database "backend/internal/database/sqlc"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type DB struct {
	Queries *database.Queries
	Pool    *pgxpool.Pool
}

func SetupDB() *DB {
	_ = godotenv.Load()
	_ = godotenv.Load("cmd/.env")
	ctx := context.Background()

	dbUser, err := getEnv("POSTGRES_USER")
	if err != nil {
		log.Fatal(err)
	}

	dbPassword, err := getEnv("POSTGRES_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}

	dbName, err := getEnv("POSTGRES_DB")
	if err != nil {
		log.Fatal(err)
	}

	dbHost, err := getEnv("DB_HOST")
	if err != nil {
		log.Fatal(err)
	}

	dbPort, err := getEnv("DB_PORT")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to create db pool: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to connect to database (%s:%s/%s): %v", dbHost, dbPort, dbName, err)
	}

	quires := database.New(pool)

	return &DB{
		Queries: quires,
		Pool:    pool,
	}
}
