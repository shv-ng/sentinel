package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ShivangSrivastava/sentinel/api"
	"github.com/ShivangSrivastava/sentinel/internal/config"
	"github.com/ShivangSrivastava/sentinel/internal/database"
	"github.com/ShivangSrivastava/sentinel/internal/logformat"

	_ "github.com/lib/pq"
)

func main() {
	// Load config
	cfg := config.Load()

	// Connect to database
	pgDB := database.ConnectPostgres(*cfg)
	// Run migrations if needed
	if ok := database.RunMigrationsIfRequired(*cfg, pgDB); ok {
		return
	}
	defer pgDB.Close()

	// Start server or other application logic
	startServer(*cfg, pgDB)
}

// startServer starts the application server
func startServer(cfg config.Config, db *sql.DB) {
	logformatRepo := logformat.NewRepo(db)
	logformatService := logformat.NewService(logformatRepo)
	logformatHandler := logformat.NewHandler(logformatService)
	r := api.NewRouter(*&logformatHandler)

	log.Println("Server is running on :", cfg.Port)
	log.Fatalf("Failed to start server: %v", http.ListenAndServe(cfg.Port, r))
}
