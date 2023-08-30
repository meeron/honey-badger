package main

import (
	"log"
	"net"

	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/pb"
	"github.com/meeron/honey-badger/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":18950")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHoneyBadgerServer(s, &server.HoneyBadgerServer{})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
