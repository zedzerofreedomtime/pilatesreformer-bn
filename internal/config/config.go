package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	AppEnv        string
	HTTPPort      string
	PostgresDSN   string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	SessionTTL    time.Duration
	CacheTTL      time.Duration
	CORSOrigins   []string
}

func Load() Config {
	return Config{
		AppEnv:        getEnv("APP_ENV", "development"),
		HTTPPort:      getEnv("HTTP_PORT", "8080"),
		PostgresDSN:   getEnv("POSTGRES_DSN", "postgres://postgres:postgres@localhost:5432/pilatesreformer?sslmode=disable"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       getEnvInt("REDIS_DB", 0),
		SessionTTL:    time.Duration(getEnvInt("SESSION_TTL_HOURS", 24)) * time.Hour,
		CacheTTL:      time.Duration(getEnvInt("CACHE_TTL_MINUTES", 5)) * time.Minute,
		CORSOrigins:   splitCSV(getEnv("CORS_ORIGINS", "http://localhost:5173")),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && strings.TrimSpace(value) != "" {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback
	}

	value, err := strconv.Atoi(raw)
	if err != nil {
		return fallback
	}

	return value
}

func splitCSV(raw string) []string {
	parts := strings.Split(raw, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}
