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

var (
	port   = 18950
	target = fmt.Sprintf("127.0.0.1:%d", port)
)

func TestDbServer(t *testing.T) {
	dbCtx := db.CreateCtx(config.BadgerConfig{
		DataDirPath: "data",
		GCPeriodMin: 60,
	})
	server := New(config.ServerConfig{
		Port:             uint16(port),
		MaxRecvMsgSizeMb: 4,
	}, dbCtx)
	conn := createConn()

	defer conn.Close()
	defer server.Stop()

	client := pb.NewDbClient(conn)
	sysClient := pb.NewSysClient(conn)

	go server.Start()

	for {
		_, err := sysClient.Ping(context.TODO(), &pb.PingRequest{})
		if err == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}

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

func createConn() *grpc.ClientConn {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)

	}

	return conn
}
