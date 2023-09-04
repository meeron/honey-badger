package server

import (
	"fmt"
	"net"

	"github.com/meeron/honey-badger/config"
	"github.com/meeron/honey-badger/logger"
	"github.com/meeron/honey-badger/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpc   *grpc.Server
	config config.ServerConfig
}

func New(c config.ServerConfig) *Server {
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * c.MaxRecvMsgSizeMb),
	}

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterHoneyBadgerServer(grpcServer, &HoneyBadgerServer{})
	reflection.Register(grpcServer)

	return &Server{
		grpc:   grpcServer,
		config: c,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return err
	}

	logger.Server().Infof("Server listening at %v", lis.Addr())
	return s.grpc.Serve(lis)
}

func (s *Server) Stop() {
	logger.Server().Infof("Stopping server...")
	s.grpc.GracefulStop()
}
