package config

import (
	"fmt"
	"log"
	"os"

	"backend/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func SetupDB() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Print("No .env file found")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("failed to connect to the database")
	}

	db.AutoMigrate(&models.User{},&models.Account{}, &models.TeamSeekPost{})

	return db
}


func CloseDB(db *gorm.DB) {
	db2, err := db.DB()


	if err != nil {
		log.Printf("failed to close the database connection")
	}

	db2.Close()
}