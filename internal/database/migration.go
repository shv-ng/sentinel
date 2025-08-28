package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// DBExecutor interface for any database that can execute SQL
type dBExecutor interface {
	Exec(ctx context.Context, query string) error
}

// PostgreSQL wrapper to implement DBExecutor
type postgresExecutor struct {
	DB *sql.DB
}

func (p *postgresExecutor) Exec(ctx context.Context, query string) error {
	_, err := p.DB.Exec(query)
	return err
}

// Generic migration runner
func RunMigrations(executor dBExecutor, migrationsPath, pattern string) error {
	files, err := filepath.Glob(filepath.Join(migrationsPath, pattern))
	if err != nil {
		return fmt.Errorf("error reading migration files: %w", err)
	}

	ctx := context.Background()
	for _, file := range files {
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("error reading %s: %w", file, err)
		}

		if err := executor.Exec(ctx, string(sqlBytes)); err != nil {
			return fmt.Errorf("migration failed on %s: %w", file, err)
		}

		log.Printf("Executed %s\n", filepath.Base(file))
	}
	return nil
}

// Helper functions for convenience
func RunPostgresMigrations(db *sql.DB, migrationsPath string) error {
	executor := &postgresExecutor{DB: db}
	return RunMigrations(executor, migrationsPath, "*_postgres.sql")
}
