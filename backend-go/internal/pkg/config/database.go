package config

import (
	database "backend/internal/database/sqlc"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)


type DB struct {
	Queries *database.Queries
	Pool *pgxpool.Pool
}


func SetupDB() *DB {
	errEnv := godotenv.Load()
	ctx := context.Background()

	if errEnv != nil {
		log.Fatal("No .env file found")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	pool, err := pgxpool.New(ctx, dsn)

	if err != nil {
		log.Fatal("failed to connect to the database")
	}

	quires := database.New(pool)

	return &DB{
		Queries: quires,
		Pool: pool,
	}
	
}