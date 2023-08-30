package main

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/meeron/honey-badger/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const BodySize = 256
const DbName = "bench"

func setValue(client pb.HoneyBadgerClient) {
	key := rand.Intn(2147483647)

	buffer := make([]byte, BodySize)
	_, err := crand.Read(buffer)
	if err != nil {
		panic(err)
	}

	res, err := client.Set(context.TODO(), &pb.SetRequest{
		Db:   DbName,
		Key:  fmt.Sprintf("%d", key),
		Data: buffer,
	})

	if err != nil {
		panic(err)
	}

	if res.Code != "ok" {
		panic(fmt.Errorf("result code is not ok: %s", res.Code))
	}
}

func getValue(client pb.HoneyBadgerClient) {
	key := rand.Intn(2147483647)

	_, err := client.Get(context.TODO(), &pb.KeyRequest{
		Db:  DbName,
		Key: fmt.Sprintf("%d", key),
	})

	if err != nil {
		panic(err)
	}
}

func BenchmarkSetValue(t *testing.B) {
	conn, err := grpc.Dial("localhost:18950", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHoneyBadgerClient(conn)

	for i := 0; i < t.N; i++ {
		setValue(client)
	}
}

func BenchmarkGetValue(t *testing.B) {
	conn, err := grpc.Dial("localhost:18950", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHoneyBadgerClient(conn)

	for i := 0; i < t.N; i++ {
		getValue(client)
	}
}
