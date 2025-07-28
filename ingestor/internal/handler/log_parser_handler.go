package handler

import (
	"context"
	"log"

	pb "github.com/ShivangSrivastava/sentinel/proto/logging"
)

type LogIngestorServer struct {
	pb.UnimplementedLogIngestorServer
}

func NewLogIngestorServer() *LogIngestorServer {
	return &LogIngestorServer{}
}

func (s *LogIngestorServer) SendLogParser(
	ctx context.Context, req *pb.LogRequest,
) (*pb.LogResponse, error) {
	log.Printf(req.GetJsonPayload())
	return &pb.LogResponse{
		Success: true,
		Message: "",
	}, nil

}
