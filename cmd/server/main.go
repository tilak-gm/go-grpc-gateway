package main

import (
	pb "github.com/tilak-gm/go-grpc-gateway/gen/go/user"
	"github.com/tilak-gm/go-grpc-gateway/internal/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	const addr = "0.0.0.0:50051"
	// create a TCP listener on the specified port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server instance
	server := grpc.NewServer()
	// register the order service with the grpc server
	pb.RegisterOrdersServer(server, user.NewService())

	// start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
