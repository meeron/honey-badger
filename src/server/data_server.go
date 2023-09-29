package server

import (
	"context"

	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/pb"
)

type DataServer struct {
	pb.UnimplementedDataServer

	dbCtx *db.DbContext
}

func (s *DataServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.EmptyResult, error) {
	db, err := s.dbCtx.GetDb(in.Db)
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

	return &pb.EmptyResult{}, nil
}

func (s *DataServer) Get(ctx context.Context, in *pb.KeyRequest) (*pb.GetResult, error) {
	db, err := s.dbCtx.GetDb(in.Db)
	if err != nil {
		return nil, err
	}

	data, hit, err := db.Get(in.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetResult{Data: data, Hit: hit}, nil
}

func (s *DataServer) GetByPrefix(ctx context.Context, in *pb.PrefixRequest) (*pb.PrefixResult, error) {
	db, err := s.dbCtx.GetDb(in.Db)
	if err != nil {
		return nil, err
	}

	data, err := db.GetByPrefix(ctx, in.Prefix)
	if err != nil {
		return nil, err
	}

	return &pb.PrefixResult{Data: data}, nil
}

func (s *DataServer) Delete(ctx context.Context, in *pb.KeyRequest) (*pb.EmptyResult, error) {
	db, err := s.dbCtx.GetDb(in.Db)
	if err != nil {
		return nil, err
	}

	if err := db.DeleteByKey(in.Key); err != nil {
		return nil, err
	}

	return &pb.EmptyResult{}, nil
}

func (s *DataServer) DeleteByPrefix(ctx context.Context, in *pb.PrefixRequest) (*pb.EmptyResult, error) {
	db, err := s.dbCtx.GetDb(in.Db)
	if err != nil {
		return nil, err
	}

	if err := db.DeleteByPrefix(in.Prefix); err != nil {
		return nil, err
	}

	return &pb.EmptyResult{}, nil
}

func (s *DataServer) SetBatch(ctx context.Context, in *pb.SetBatchRequest) (*pb.EmptyResult, error) {
	db, err := s.dbCtx.GetDb(in.Db)
	if err != nil {
		return nil, err
	}

	if err := db.SetBatch(in.Data); err != nil {
		return nil, err
	}

	return &pb.EmptyResult{}, nil
}

func (s *DataServer) GetDataStream(in *pb.DataStreamRequest, stream pb.Data_GetDataStreamServer) error {
	db, err := s.dbCtx.GetDb(in.Db)
	if err != nil {
		return err
	}

	if in.Prefix == nil {
		return nil
	}

	data, err := db.GetByPrefix(stream.Context(), *in.Prefix)
	if err != nil {
		return err
	}

	for key, itm := range data {
		stream.Send(&pb.DataItem{
			Key:  key,
			Data: itm,
		})
	}

	return nil
}
