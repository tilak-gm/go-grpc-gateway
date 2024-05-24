package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/grpclog"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/tilak-gm/go-grpc-gateway/gen/go/user"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterUsersHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		log.Printf("failed to register the user server: %v", err)
		grpclog.Fatal(err)
	}

	// start listening to requests from the gateway server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Printf("gateway server closed abruptly: %v", err)
		grpclog.Fatal(err)
	} else {
		fmt.Println("API gateway server is running on :8080")
	}
}
