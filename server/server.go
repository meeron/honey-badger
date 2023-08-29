package server

import (
	"context"

	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/pb"
)

type HoneyBadgerServer struct {
	pb.UnimplementedHoneyBadgerServer
}

func (s *HoneyBadgerServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.Result, error) {
	db, err := db.Get(in.Db)
	if err != nil {
		return nil, err
	}

	var ttl uint = 0
	if in.Ttl != nil {
		ttl = uint(*in.Ttl)
	}

	err = db.Set(in.Key, in.Data, 0, ttl)
	if err != nil {
		return nil, err
	}

	return &pb.Result{Code: "ok"}, nil
}

func (s *HoneyBadgerServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResult, error) {
	db, err := db.Get(in.Db)
	if err != nil {
		return nil, err
	}

	data, hit, err := db.Get(in.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetResult{Data: data, Hit: hit}, nil
}
