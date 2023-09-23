package server

import (
	"context"

	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/pb"
)

type DbServer struct {
	pb.UnimplementedDbServer

	dbCtx *db.DbContext
}

func (s *DbServer) Create(ctx context.Context, in *pb.CreateDbRequest) (*pb.Result, error) {
	_, err := s.dbCtx.CreateDb(in.Name, in.InMemory)
	if err != nil {
		return nil, err
	}

	return &pb.Result{Code: "ok"}, nil
}

func (s *DbServer) Drop(ctx context.Context, in *pb.DropDbRequest) (*pb.Result, error) {
	if err := s.dbCtx.DropDb(in.Name); err != nil {
		return nil, err
	}

	return &pb.Result{Code: "ok"}, nil
}
