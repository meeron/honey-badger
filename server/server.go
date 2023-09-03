package server

import (
	"context"
	"fmt"
	"net"

	"github.com/meeron/honey-badger/config"
	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/logger"
	"github.com/meeron/honey-badger/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	err = db.Set(in.Key, in.Data, ttl)
	if err != nil {
		return nil, err
	}

	return &pb.Result{Code: "ok"}, nil
}

func (s *HoneyBadgerServer) Get(ctx context.Context, in *pb.KeyRequest) (*pb.GetResult, error) {
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

func (s *HoneyBadgerServer) GetByPrefix(ctx context.Context, in *pb.PrefixRequest) (*pb.PrefixResult, error) {
	db, err := db.Get(in.Db)
	if err != nil {
		return nil, err
	}

	data, err := db.GetByPrefix(ctx, in.Prefix)
	if err != nil {
		return nil, err
	}

	return &pb.PrefixResult{Data: data}, nil
}

func (s *HoneyBadgerServer) Delete(ctx context.Context, in *pb.KeyRequest) (*pb.Result, error) {
	db, err := db.Get(in.Db)
	if err != nil {
		return nil, err
	}

	if err := db.DeleteByKey(in.Key); err != nil {
		return nil, err
	}

	return &pb.Result{Code: "ok"}, nil
}

func (s *HoneyBadgerServer) DeleteByPrefix(ctx context.Context, in *pb.PrefixRequest) (*pb.Result, error) {
	db, err := db.Get(in.Db)
	if err != nil {
		return nil, err
	}

	if err := db.DeleteByPrefix(in.Prefix); err != nil {
		return nil, err
	}

	return &pb.Result{Code: "ok"}, nil
}

func (s *HoneyBadgerServer) SetBatch(ctx context.Context, in *pb.SetBatchRequest) (*pb.Result, error) {
	db, err := db.Get(in.Db)
	if err != nil {
		return nil, err
	}

	if err := db.SetBatch(in.Data); err != nil {
		return nil, err
	}

	return &pb.Result{Code: "ok"}, nil
}

func Run() error {
	config := config.Get().Server

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * config.MaxRecvMsgSizeMb),
	}

	s := grpc.NewServer(opts...)

	pb.RegisterHoneyBadgerServer(s, &HoneyBadgerServer{})
	reflection.Register(s)

	logger.Server().Infof("Server listening at %v", lis.Addr())
	return s.Serve(lis)
}
