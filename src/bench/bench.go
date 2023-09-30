package bench

import (
	"context"
	"fmt"
	"io"
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
	DbName          = "bench_v2"
	PayloadSize     = 256
	NumGoProc       = 20
	BatchItemPrefix = "batch-item"
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

func benchSendStream(target string) {
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
		stream, err := client.CreateSendStream(context.TODO())
		if err != nil {
			panic(err)
		}

		// First message must contain db name
		dbReq := &pb.SendStreamReq{
			Db: DbName,
		}
		if err := stream.Send(dbReq); err != nil {
			panic(err)
		}

		start := time.Now()
		for j := 0; j < batchIts[i]; j++ {
			req := &pb.SendStreamReq{
				Item: &pb.DataItem{
					Key:  fmt.Sprintf("%s-%d-%d", BatchItemPrefix, i, j),
					Data: make([]byte, PayloadSize),
				},
			}

			if err := stream.Send(req); err != nil {
				panic(err)
			}
		}

		_, err = stream.CloseAndRecv()
		if err != nil {
			panic(err)
		}

		fmt.Printf("SendWithStream_%d: %s\n", batchIts[i], time.Since(start))
	}
}

func benchReadStream(target string) {
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
		prefix := fmt.Sprintf("%s-%d-", BatchItemPrefix, i)

		start := time.Now()
		stream, errGet := client.CreateReadStream(context.TODO(), &pb.ReadStreamReq{
			Db:     DbName,
			Prefix: &prefix,
		})
		if errGet != nil {
			panic(errGet)
		}

		for {
			_, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
		}

		fmt.Printf("ReadWithStream_%d: %s\n", batchIts[i], time.Since(start))
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
	benchSendStream(target)
	benchReadStream(target)
}
