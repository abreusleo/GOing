package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
}

func LoadConfig() Config {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	envPath := filepath.Join(wd, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	} else {
		log.Printf(".env file loaded successfully from %s", envPath)
	}

	cfg := Config{
		DBUser:     getenv("DB_USER"),
		DBPassword: getenv("DB_PASSWORD"),
		DBName:     getenv("DB_NAME"),
		DBHost:     getenv("DB_HOST"),
	}

	return cfg
}

func getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}
