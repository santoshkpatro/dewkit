package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func loadEnv() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment")
		}
	})
}

// GetEnv returns the environment variable value or empty string
func GetEnv(key string) string {
	loadEnv()
	return os.Getenv(key)
}

// GetEnvDefault returns env var or default value
func GetEnvDefault(key, defaultVal string) string {
	loadEnv()
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
