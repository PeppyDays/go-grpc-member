package rpc

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

// TODO: Modify to bind port from external
type ServerOption func(*Server)

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

type Server struct {
	grpc        *grpc.Server
	port        int
	baseHandler *BaseHandler
}

func (s *Server) Server() *grpc.Server {
	return s.grpc
}

func (s *Server) Serve() error {
	address := fmt.Sprintf("0.0.0.0:%d", s.port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to create gRPC listener: %v", err)
	}
	slog.Info("gPRC server start to listen")
	return s.grpc.Serve(listener)
}

func NewServer(h *BaseHandler, opts ...ServerOption) (*Server, error) {
	server := &Server{
		port:        50051,
		grpc:        grpc.NewServer(),
		baseHandler: h,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server, nil
}
