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
	pb.UnimplementedOrdersServer
}

func NewService() *Service {
	return &Service{}
}

func (u *Service) GetUser(ctx context.Context, req *pb.PayloadWithUserId) (*pb.PayloadWithSingleUser, error) {
	log.Printf("Received GET user request with payload: %v", req)

	user, err := u.GetUser(context.TODO(), req)
	if err != nil {
		log.Fatal("UserService.GetUser err: ", err)
	}
	return user, nil
}
