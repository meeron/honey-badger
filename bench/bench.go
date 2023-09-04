package bench

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/meeron/honey-badger/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	getSetIts = []int{
		10_000,
		30_000,
		50_000,
	}
	batchIts = []int{
		50_000,
		100_000,
		300_000,
	}
)

const (
	DbName      = "bench"
	PayloadSize = 256
)

func benchSet(target string) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHoneyBadgerClient(conn)

	payload := make([]byte, PayloadSize)

	for i := 0; i < len(getSetIts); i++ {
		start := time.Now()

		for j := 0; j < getSetIts[i]; j++ {
			_, err := client.Set(context.TODO(), &pb.SetRequest{
				Db:   DbName,
				Key:  fmt.Sprintf("bench-test-%d", j),
				Data: payload,
			})
			if err != nil {
				panic(err)
			}
		}

		fmt.Printf("Set_%d: %s\n", getSetIts[i], time.Since(start))
	}
}

func benchGet(target string) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHoneyBadgerClient(conn)

	for i := 0; i < len(getSetIts); i++ {
		start := time.Now()

		for j := 0; j < getSetIts[i]; j++ {
			_, err := client.Get(context.TODO(), &pb.KeyRequest{
				Db:  DbName,
				Key: fmt.Sprintf("bench-test-%d", j),
			})
			if err != nil {
				panic(err)
			}
		}

		fmt.Printf("Get_%d: %s\n", getSetIts[i], time.Since(start))
	}
}

func benchSetBatch(target string) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHoneyBadgerClient(conn)

	for i := 0; i < len(batchIts); i++ {
		data := make(map[string][]byte)

		for j := 0; j < batchIts[i]; j++ {
			data[fmt.Sprintf("batch-%d", j)] = make([]byte, PayloadSize)
		}

		start := time.Now()

		_, err := client.SetBatch(context.TODO(), &pb.SetBatchRequest{
			Db:   DbName,
			Data: data,
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("SetBatch_%d: %s\n", batchIts[i], time.Since(start))
	}
}

func Run(target string) {
	benchSet(target)
	benchGet(target)
	benchSetBatch(target)
}
