package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/youcan-jpn/janken_go_grpc/pb"
	"github.com/youcan-jpn/janken_go_grpc/service"
)

func main() {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		log.Fatalf("AAA", err)
	}

	server := grpc.NewServer()

	pb.RegisterRockPaperScissorsServiceServer(server, service.NewRockPaperScissorsService())
	reflection.Register(server)
	server.Serve(listenPort)
}
