package user

import (
	"context"
	pb "github.com/tilak-gm/go-grpc-gateway/gen/go/user"
	"log"
)

// Service should implement the OrdersServer interface generated from grpc.
//
// UnimplementedOrdersServer must be embedded to have forwarded compatible implementations.
type Service struct {
	pb.UnimplementedUsersServer
}

func NewService() *Service {
	return &Service{}
}

func (u *Service) GetUser(ctx context.Context, req *pb.PayloadWithUserId) (*pb.PayloadWithSingleUser, error) {
	log.Printf("Received GET user request with payload: %v", req)

	// TODO: need to integrate with DB for CRUD operations on Users
	// returning dummy for now,
	user := &pb.User{
		UserId: 1,
		Name:   "tilak",
		Age:    26,
	}
	return &pb.PayloadWithSingleUser{User: user}, nil
}
