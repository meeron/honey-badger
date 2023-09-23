package server

import (
	"context"

	"github.com/meeron/honey-badger/pb"
)

type SysServer struct {
	pb.UnimplementedSysServer
}

func (s *SysServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.Result, error) {
	return &pb.Result{Code: "pong"}, nil
}
