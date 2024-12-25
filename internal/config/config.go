package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config структура, содержащая конфигурационные параметры сервера
type Config struct {
	ServerPort string
}

// LoadConfig загружает параметры для запуска сервера
func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
		return nil
	}

	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		port = ":8080"
	}

	return &Config{
		ServerPort: port,
	}
}
