package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN string

	ClickHouseHost     string
	ClickHousePort     string
	ClickHouseDatabase string
	ClickHouseUsername string
	ClickHousePassword string

	RunMigrations string

	Port string
}

func Load() *Config {
	// Load .env if present
	_ = godotenv.Load()

	cfg := &Config{
		PostgresDSN: getEnv("POSTGRES_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),

		ClickHouseHost:     getEnv("CLICKHOUSE_HOST", "localhost"),
		ClickHousePort:     getEnv("CLICKHOUSE_TCP_PORT", "9000"),
		ClickHouseDatabase: getEnv("CLICKHOUSE_DB", "default"),
		ClickHouseUsername: getEnv("CLICKHOUSE_USER", "default"),
		ClickHousePassword: getEnv("CLICKHOUSE_PASSWORD", ""),
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
