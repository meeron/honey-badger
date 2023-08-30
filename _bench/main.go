package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/meeron/honey-badger/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var its = []int{
	10_000,
	30_000,
	50_000,
}
var dbName = "bench"

func benchSet() {
	conn, err := grpc.Dial("localhost:18950", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHoneyBadgerClient(conn)

	payload := make([]byte, 256)

	for i := 0; i < len(its); i++ {
		start := time.Now()

		for j := 0; j < its[i]; j++ {
			_, err := client.Set(context.TODO(), &pb.SetRequest{
				Db:   dbName,
				Key:  fmt.Sprintf("bench-test-%d", j),
				Data: payload,
			})
			if err != nil {
				panic(err)
			}
		}

		fmt.Printf("Set_%d: %s\n", its[i], time.Since(start))
	}
}

func benchGet() {
	conn, err := grpc.Dial("localhost:18950", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHoneyBadgerClient(conn)

	for i := 0; i < len(its); i++ {
		start := time.Now()

		for j := 0; j < its[i]; j++ {
			_, err := client.Get(context.TODO(), &pb.KeyRequest{
				Db:  dbName,
				Key: fmt.Sprintf("bench-test-%d", j),
			})
			if err != nil {
				panic(err)
			}
		}

		fmt.Printf("Get_%d: %s\n", its[i], time.Since(start))
	}
}

func main() {
	benchSet()
	benchGet()
}
