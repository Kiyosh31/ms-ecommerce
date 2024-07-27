package handler

import (
	"context"
	"log"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/user-service/service"

	"google.golang.org/grpc"
)

type UserServiceGrpcHandler struct {
	userPb.UnimplementedUserServiceServer
	service service.UserService
}

func NewGrpcUserServiceHandler(grpcServer *grpc.Server, service service.UserService) {
	handler := &UserServiceGrpcHandler{
		service: service,
	}

	userPb.RegisterUserServiceServer(grpcServer, handler)
}

func (c *UserServiceGrpcHandler) CreateUser(ctx context.Context, in *userPb.CreateUserRequest) (*userPb.Response, error) {
	log.Printf("Create user received request! %v", in)

	res := &userPb.Response{
		Message: "User created successfully!",
		User: &userPb.User{
			Id:       "1",
			Name:     "david",
			LastName: "Garcia",
			Birth:    "123456",
			Cards: &userPb.CardList{
				Cards: []*userPb.Card{
					{
						Id:         "1",
						Number:     1234567890,
						Cvv:        123,
						Expiration: "1234",
						Default:    true,
					},
				},
			},
			Addresses: &userPb.AddressList{
				Address: []*userPb.Address{
					{
						Id:      "1",
						Name:    "street",
						ZipCode: 1234,
						Default: true,
					},
				},
			},
			Email:    "email",
			Password: "password",
		},
	}

	return res, nil
}
