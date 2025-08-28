package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN string

	RunMigrations string

	Port string
}

func Load() *Config {
	// Load .env if present
	_ = godotenv.Load()

	cfg := &Config{
		PostgresDSN: getEnv("POSTGRES_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),

		RunMigrations: getEnv("RUN_MIGRATIONS", "false"),

		Port: getEnv("PORT", "8080"),
	}

	return cfg
}

// Helper function to get environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
