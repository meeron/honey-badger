package server

import (
	"context"

	"github.com/meeron/honey-badger/pb"
)

type HoneyBadgerServer struct {
	pb.UnimplementedHoneyBadgerServer
}

func (s *HoneyBadgerServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.Result, error) {
	//log.Printf("Set: %v", in)
	return &pb.Result{Code: "ok"}, nil
}

func (s *HoneyBadgerServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResult, error) {
	//log.Printf("Get: %v", in)
	return &pb.GetResult{Data: make([]byte, 0)}, nil
}
