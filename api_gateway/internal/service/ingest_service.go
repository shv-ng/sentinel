package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ShivangSrivastava/sentinel/api_gateway/internal/utils"
	"google.golang.org/grpc"

	pb "github.com/ShivangSrivastava/sentinel/proto/logging"
)

type IngestorService interface {
	SendLogParser(rawJSON string) (*pb.LogResponse, error)
	Close() error
}
type ingestorService struct {
	client pb.LogIngestorClient
	conn   *grpc.ClientConn
}

func NewIngestorService(address string) (IngestorService, error) {
	conn, err := utils.ConnectWithRetry(address, 10)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ingestor at %s: %w", address, err)
	}

	client := pb.NewLogIngestorClient(conn)
	return &ingestorService{
		client: client,
		conn:   conn,
	}, nil
}

func (s *ingestorService) SendLogParser(rawJSON string) (*pb.LogResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logReq := &pb.LogRequest{
		JsonPayload: rawJSON,
	}

	response, err := s.client.SendLogParser(ctx, logReq)
	if err != nil {
		return nil, fmt.Errorf("SendLogParser failed: %w", err)
	}

	return response, nil
}

func (s *ingestorService) Close() error {
	if s.conn != nil {
		return s.conn.Close()
	}
	return nil
}
