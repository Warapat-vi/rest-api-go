package serverConfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DataBaseUrl string
	Environment string
	LogLevel    string
}

func LoadConfig() (*Config, error) {
	if error := godotenv.Load(); error != nil {
		return nil, fmt.Errorf("Error loading file: %v", error)
	}
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DataBaseUrl: getEnv("DATABASE_URL", "postgres"),
		Environment: getEnv("ENVIRONMENT", "local"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
