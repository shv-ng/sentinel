package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/shv-ng/sentinel/api"
	"github.com/shv-ng/sentinel/internal/database"
	"github.com/shv-ng/sentinel/internal/logformat"

	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	loadEnvVariables()

	// Connect to database
	pgDB := database.ConnectToPostgres()
	// Run migrations if needed
	if ok := database.RunMigrationsIfRequired(pgDB); ok {
		return
	}
	defer pgDB.Close()

	// Start server or other application logic
	startServer(pgDB)
}

// loadEnvVariables loads environment variables from .env file
func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
	log.Println("Environment variables loaded successfully")
}

// startServer starts the application server
func startServer(db *sql.DB) {
	logformatRepo := logformat.NewRepo(db)
	logformatService := logformat.NewService(logformatRepo)
	logformatHandler := logformat.NewHandler(logformatService)
	r := api.NewRouter(*&logformatHandler)

	log.Println("Server is running on :8080")
	log.Fatalf("Failed to start server: %v", http.ListenAndServe(":8080", r))
}
