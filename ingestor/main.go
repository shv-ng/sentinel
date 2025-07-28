package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/ShivangSrivastava/sentinel/ingestor/internal/database"
	"github.com/ShivangSrivastava/sentinel/ingestor/internal/handler"
	"github.com/ShivangSrivastava/sentinel/ingestor/internal/repo"
	"github.com/ShivangSrivastava/sentinel/ingestor/internal/service"
	pb "github.com/ShivangSrivastava/sentinel/proto/logging"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
	pgURL := os.Getenv("POSTGRES_URL")
	pgDB := database.ConnectToPostgres(pgURL)
	if ok := database.RunMigrationsIfRequired(pgDB); ok {
		return
	}
	defer pgDB.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()
	logParserRepo := repo.NewLogParserRepo(pgDB)
	logParserservice := service.NewLogParserService(logParserRepo)
	logParserHandler := handler.NewLogIngestorServer(logParserservice)

	grpcServer := grpc.NewServer()

	pb.RegisterLogIngestorServer(grpcServer, logParserHandler)
	log.Println("LogIngestor gRPC server running on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
