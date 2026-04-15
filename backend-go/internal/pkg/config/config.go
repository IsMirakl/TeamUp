package config

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type EnvConfig struct {
	JWT_SECRET string
	REFRESH_SECRET string
}


type Config struct {
	SECRET_KEY EnvConfig
	log *logrus.Logger
}

func New(log *logrus.Logger) *Config{
	return &Config{
		SECRET_KEY: EnvConfig{
			JWT_SECRET: getEnv("SECRET_KEY"),
			REFRESH_SECRET: getEnv("REFRESH_TOKEN_KEY"),
		},
		log: log,
	}
}


func getEnv(key string) string {
	exists, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("Environment variable not set")
	}
	return exists
}