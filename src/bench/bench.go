package bench

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/meeron/honey-badger/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	getSetIts = []int{
		30_000,
		50_000,
		100_000,
	}
	batchIts = []int{
		100_000,
		300_000,
		500_000,
	}
)

const (
	DbName      = "bench"
	PayloadSize = 256
	NumGoProc   = 20
)

func benchSet(target string) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataClient(conn)

	payload := make([]byte, PayloadSize)

	fmt.Println("")
	fmt.Printf("payload size: %d bytes\n", PayloadSize)
	fmt.Printf("num goroutines: %d\n", NumGoProc)

	for i := 0; i < len(getSetIts); i++ {
		limiter := make(chan int, NumGoProc)
		wg := new(sync.WaitGroup)
		wg.Add(getSetIts[i])

		start := time.Now()
		for j := 0; j < getSetIts[i]; j++ {
			limiter <- j
			go sendSet(j, client, payload, limiter, wg)
		}
		wg.Wait()
		fmt.Printf("Set_%d: %s\n", getSetIts[i], time.Since(start))
	}
}

func benchGet(target string) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataClient(conn)

	fmt.Println("")
	fmt.Printf("payload size: %d bytes\n", PayloadSize)
	fmt.Printf("num goroutines: %d\n", NumGoProc)

	for i := 0; i < len(getSetIts); i++ {
		limiter := make(chan int, NumGoProc)
		wg := new(sync.WaitGroup)
		wg.Add(getSetIts[i])

		start := time.Now()
		for j := 0; j < getSetIts[i]; j++ {
			limiter <- j
			go sendGet(j, client, limiter, wg)
		}
		wg.Wait()
		fmt.Printf("Get_%d: %s\n", getSetIts[i], time.Since(start))
	}
}

func benchSetBatch(target string) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataClient(conn)

	fmt.Println("")
	fmt.Printf("payload size: %d bytes\n", PayloadSize)
	fmt.Printf("num goroutines: %d\n", 1)

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

func sendSet(index int, client pb.DataClient, payload []byte, limiter <-chan int, wg *sync.WaitGroup) {
	_, err := client.Set(context.TODO(), &pb.SetRequest{
		Db:   DbName,
		Key:  fmt.Sprintf("bench-test-%d", index),
		Data: payload,
	})
	if err != nil {
		panic(err)
	}
	wg.Done()
	<-limiter
}

func sendGet(index int, client pb.DataClient, limiter <-chan int, wg *sync.WaitGroup) {
	_, err := client.Get(context.TODO(), &pb.KeyRequest{
		Db:  DbName,
		Key: fmt.Sprintf("bench-test-%d", index),
	})
	if err != nil {
		panic(err)
	}
	wg.Done()
	<-limiter
}

func Run(target string) {
	fmt.Printf("os: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("cpus: %d\n", runtime.NumCPU())

	benchSet(target)
	benchGet(target)
	benchSetBatch(target)
}
