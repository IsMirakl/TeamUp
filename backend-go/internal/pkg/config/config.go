package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type EnvConfig struct {
	JWT_SECRET     string
	REFRESH_SECRET string
}

type Config struct {
	SECRET_KEY EnvConfig
	log        *logrus.Logger
}

func New(log *logrus.Logger) (*Config, error) {
	// Load a local .env file when present (dev convenience).
	// Ignore errors so production relies on real environment variables.
	_ = godotenv.Load()
	_ = godotenv.Load("cmd/.env")

	jwtSecret, err := getEnv("SECRET_KEY")
	if err != nil {
		return nil, err
	}

	refreshSecret, err := getEnv("REFRESH_TOKEN_KEY")
	if err != nil {
		return nil, err
	}

	return &Config{
		SECRET_KEY: EnvConfig{
			JWT_SECRET:     jwtSecret,
			REFRESH_SECRET: refreshSecret,
		},
		log: log,
	}, nil
}

func getEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return value, nil
}
