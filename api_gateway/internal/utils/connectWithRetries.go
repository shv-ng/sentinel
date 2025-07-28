package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ShivangSrivastava/sentinel/proto/logging"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func ConnectWithRetry(address string, maxRetries int) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error

	for i := range maxRetries {
		log.Printf("Connection attempt %d to %s", i+1, address)

		conn, err = grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to create client: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		client := pb.NewLogIngestorClient(conn)

		_, testErr := client.SendLogParser(ctx, &pb.LogRequest{
			JsonPayload: `{}`,
		})
		cancel()

		if testErr == nil {
			log.Printf("Successfully connected to %s", address)
			return conn, nil
		}

		// Check if it's a connection error vs application error
		if status.Code(testErr) != codes.Unavailable {
			log.Printf("Connected but got application error: %v", testErr)
			return conn, nil // Connection works, application error is fine
		}

		log.Printf("Connection test failed: %v", testErr)
		conn.Close()
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect after %d attempts: %v", maxRetries, err)
}
