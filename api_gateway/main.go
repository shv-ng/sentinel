package main

import (
	"log"
	"net/http"

	"github.com/ShivangSrivastava/sentinel/api_gateway/internal/handler"
	"github.com/ShivangSrivastava/sentinel/api_gateway/internal/router"
	"github.com/ShivangSrivastava/sentinel/api_gateway/internal/service"
)

func main() {
	ingestService, err := service.NewIngestorService("ingestor:50051")
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer ingestService.Close()
	ingestHandler := handler.NewIngestHandler(ingestService)

	r := router.NewRouter(ingestHandler)
	log.Println("Server is running on :8080")
	log.Fatalf("Failed to start server: %v\n", http.ListenAndServe(":8080", r))
}
