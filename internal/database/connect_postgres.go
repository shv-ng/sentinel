package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/ShivangSrivastava/sentinel/internal/config"
)

// establishes database connection and verifies connectivity
func ConnectPostgres(cfg config.Config) *sql.DB {

	dbURL := cfg.PostgresDSN
	if dbURL == "" {
		log.Fatalln("POSTGRES_URL not found in environment")
	}
	const maxRetries = 5
	const baseDelay = time.Second

	var db *sql.DB
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		log.Printf("Attempting database connection (attempt %d/%d)", attempt, maxRetries)

		db, err = sql.Open("postgres", dbURL)
		if err != nil {
			log.Printf("Failed to open database connection on attempt %d: %v", attempt, err)
			if attempt == maxRetries {
				log.Fatalf("Failed to open database connection after %d attempts: %v", maxRetries, err)
			}
			time.Sleep(baseDelay * time.Duration(attempt)) // Exponential backoff
			continue
		}

		// Verify database connectivity
		if err = db.Ping(); err != nil {
			log.Printf("Database not reachable on attempt %d: %v", attempt, err)
			db.Close() // Close the connection before retrying
			if attempt == maxRetries {
				log.Fatalf("Database not reachable after %d attempts: %v", maxRetries, err)
			}
			time.Sleep(baseDelay * time.Duration(attempt)) // Exponential backoff
			continue
		}

		log.Println("Database connected successfully")
		return db
	}
	return nil
}
