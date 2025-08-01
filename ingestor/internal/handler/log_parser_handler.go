package handler

import (
	"context"
	"fmt"

	"github.com/ShivangSrivastava/sentinel/ingestor/internal/service"
	pb "github.com/ShivangSrivastava/sentinel/proto/logging"
)

type LogIngestorServer struct {
	pb.UnimplementedLogIngestorServer
	service service.LogParserService
}

func NewLogIngestorServer(s service.LogParserService) *LogIngestorServer {
	return &LogIngestorServer{
		service: s,
	}
}

func (s *LogIngestorServer) SendLogParser(
	ctx context.Context, req *pb.ParserFormatRequest,
) (*pb.ParserFormatResponse, error) {
	err := s.service.CreateLogFormat(req.GetJsonPayload())
	if err != nil {
		return &pb.ParserFormatResponse{
			Success: false,
			Message: fmt.Sprintf("%v", err),
		}, err
	}
	return &pb.ParserFormatResponse{
		Success: true,
		Message: "Successfully added",
	}, nil

}
