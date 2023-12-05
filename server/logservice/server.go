package logservice

import (
	"context"

	"github.com/1x-eng/struct_laas/gen/logger"
	"google.golang.org/grpc"
)

type logServiceServer struct {
	logger.UnimplementedLoggerServiceServer
}

func (s *logServiceServer) Log(ctx context.Context, in *logger.LogMessage) (*logger.LogResponse, error) {
	// TODO: get this to log aggregator of choice? for now, dump it to std out.
	ProcessLog(in)
	return &logger.LogResponse{Success: true}, nil
}

func Register(s *grpc.Server) {
	logger.RegisterLoggerServiceServer(s, &logServiceServer{})
}
