package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/meeron/honey-badger/config"
	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDbServer(t *testing.T) {
	conn, server := startServer()
	defer conn.Close()
	defer server.Stop()

	client := pb.NewDbClient(conn)

	t.Run("should call create database", func(t *testing.T) {
		res, err := client.Create(context.TODO(), &pb.CreateDbRequest{
			Name:     "test-db",
			InMemory: true,
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))

		if res != nil {
			assert.Equal(t, "ok", res.Code)
		}
	})

	t.Run("should call drop database", func(t *testing.T) {
		res, err := client.Drop(context.TODO(), &pb.DropDbRequest{
			Name: "test-db",
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))

		if res != nil {
			assert.Equal(t, "ok", res.Code)
		}
	})
}

func TestDataServer(t *testing.T) {
	conn, server := startServer()
	defer conn.Close()
	defer server.Stop()

	client := pb.NewDataClient(conn)

	t.Run("should call set", func(t *testing.T) {
		res, err := client.Set(context.TODO(), &pb.SetRequest{
			Db:   "test-db",
			Key:  "test-key",
			Data: []byte("test"),
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))

		if res != nil {
			assert.Equal(t, "ok", res.Code)
		}
	})

	t.Run("should call get", func(t *testing.T) {
		_, err := client.Get(context.TODO(), &pb.KeyRequest{
			Db:  "test-db",
			Key: "test-key",
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))
	})

	t.Run("should call get by prefix", func(t *testing.T) {
		_, err := client.GetByPrefix(context.TODO(), &pb.PrefixRequest{
			Db:     "test-db",
			Prefix: "test-",
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))
	})

	t.Run("should call delete", func(t *testing.T) {
		res, err := client.Delete(context.TODO(), &pb.KeyRequest{
			Db:  "test-db",
			Key: "test-test",
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))

		if res != nil {
			assert.Equal(t, "ok", res.Code)
		}
	})

	t.Run("should call delete by prefix", func(t *testing.T) {
		res, err := client.DeleteByPrefix(context.TODO(), &pb.PrefixRequest{
			Db:     "test-db",
			Prefix: "test-",
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))

		if res != nil {
			assert.Equal(t, "ok", res.Code)
		}
	})

	t.Run("should call set batch", func(t *testing.T) {
		res, err := client.SetBatch(context.TODO(), &pb.SetBatchRequest{
			Db:   "test-db",
			Data: make(map[string][]byte),
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))

		if res != nil {
			assert.Equal(t, "ok", res.Code)
		}
	})
}

func startServer() (*grpc.ClientConn, *Server) {
	port := 18950
	target := fmt.Sprintf("127.0.0.1:%d", port)

	dbCtx := db.CreateCtx(config.BadgerConfig{
		DataDirPath: "data",
		GCPeriodMin: 60,
	})
	server := New(config.ServerConfig{
		Port:             uint16(port),
		MaxRecvMsgSizeMb: 4,
	}, dbCtx)

	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)

	}

	sysClient := pb.NewSysClient(conn)

	go server.Start()

	for {
		_, err := sysClient.Ping(context.TODO(), &pb.PingRequest{})
		if err == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}

	return conn, server
}
