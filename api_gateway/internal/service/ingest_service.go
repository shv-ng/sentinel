package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ShivangSrivastava/sentinel/api_gateway/internal/utils"
	"google.golang.org/grpc"

	pb "github.com/ShivangSrivastava/sentinel/proto/logging"
)

type IngestorClient struct {
	client pb.LogIngestorClient
	conn   *grpc.ClientConn
}

func NewIngestorService(address string) (*IngestorClient, error) {
	conn, err := utils.ConnectWithRetry(address, 10)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ingestor at %s: %w", address, err)
	}

	client := pb.NewLogIngestorClient(conn)
	return &IngestorClient{
		client: client,
		conn:   conn,
	}, nil
}

func (s *IngestorClient) SendLogParser(rawJSON string) (*pb.LogResponse, error) {
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

func (c *IngestorClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
