package config

import (
	"log"
	"os"
)

type EnvConfig struct {
	JWT_SECRET string
}

type Config struct {
	SECRET_KEY EnvConfig
}

func New() *Config{
	return &Config{
		SECRET_KEY: EnvConfig{
			JWT_SECRET: getEnv("SECRET_KEY"),
		},
	}
}


func getEnv(key string) string {
	exists, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("Environment variable not set")
	}
	return exists
}