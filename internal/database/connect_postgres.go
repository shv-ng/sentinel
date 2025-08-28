package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/ShivangSrivastava/sentinel/internal/config"
)

// establishes database connection and verifies connectivity
func ConnectPostgres(cfg config.Config) *sql.DB {

	const maxRetries = 5
	const baseDelay = time.Second

	var db *sql.DB
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		log.Printf("Attempting postgres connection (attempt %d/%d)", attempt, maxRetries)

		db, err = sql.Open("postgres", cfg.PostgresDSN)
		if err != nil {
			log.Printf("Failed to open postgres connection on attempt %d: %v", attempt, err)
			if attempt == maxRetries {
				log.Fatalf("Failed to open postgres connection after %d attempts: %v", maxRetries, err)
			}
			time.Sleep(baseDelay * time.Duration(attempt)) // Exponential backoff
			continue
		}

		// Verify database connectivity
		if err = db.Ping(); err != nil {
			log.Printf("Postgres not reachable on attempt %d: %v", attempt, err)
			db.Close() // Close the connection before retrying
			if attempt == maxRetries {
				log.Fatalf("Postgres not reachable after %d attempts: %v", maxRetries, err)
			}
			time.Sleep(baseDelay * time.Duration(attempt)) // Exponential backoff
			continue
		}

		log.Println("Postgres connected successfully")
		return db
	}
	return nil
}
