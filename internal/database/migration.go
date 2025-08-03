package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func runMigrations(db *sql.DB, migrationsPath string) error {
	files, err := filepath.Glob(filepath.Join(migrationsPath, "*.sql"))
	if err != nil {
		return fmt.Errorf("error reading migration files: %w", err)
	}

	for _, file := range files {
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("error reading %s: %w", file, err)
		}
		if _, err := db.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("migration failed on %s: %w", file, err)
		}
		fmt.Printf("Executed %s\n", filepath.Base(file))
	}
	return nil
}

// runMigrationsIfRequired runs database migrations when RUN_MIGRATIONS=true
func RunMigrationsIfRequired(db *sql.DB) bool {
	if os.Getenv("RUN_MIGRATIONS") != "true" {
		return false
	}

	log.Println("Running database migrations...")
	if err := runMigrations(db, "migrations/"); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")
	return true
}
