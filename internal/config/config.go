package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort        string
	DB_HOST           string
	DB_PORT           string
	DB_USER           string
	DB_PASSWORD       string
	DB_NAME           string
	REDIS_HOST        string
	REDIS_PORT        string
	MINIO_ENDPOINT    string
	MINIO_ACCESS_KEY  string
	MINIO_SECRET_KEY  string
	MINIO_BUCKET_NAME string
	JWT_SECRET        string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	config := &Config{
		ServerPort:        getEnv("SERVER_PORT", "8080"),
		DB_HOST:           getEnv("DB_HOST", "localhost"),
		DB_PORT:           getEnv("DB_PORT", "5432"),
		DB_USER:           getEnv("DB_USER", "postgres"),
		DB_PASSWORD:       getEnv("DB_PASSWORD", "root"),
		DB_NAME:           getEnv("DB_NAME", "yalp_db"),
		REDIS_HOST:        getEnv("REDIS_HOST", "localhost"),
		REDIS_PORT:        getEnv("REDIS_PORT", "6379"),
		MINIO_ENDPOINT:    getEnv("MINIO_ENDPOINT", "localhost:9002"),
		MINIO_ACCESS_KEY:  getEnv("MINIO_ACCESS_KEY", "minioadmin"),
		MINIO_SECRET_KEY:  getEnv("MINIO_SECRET_KEY", "miniopassword"),
		MINIO_BUCKET_NAME: getEnv("MINIO_BUCKET_NAME", "images"),
		JWT_SECRET:        getEnv("JWT_SECRET", "cd46911f94334c65301d4086f14f0eb29ef7a4ecb3aba5ced1e7ec9760a4356c74822dea0bbe2ff636a3bd1483fd9472b7f016ddffe91603f0b32854143d173c"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
