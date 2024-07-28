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

func (s *UserServiceGrpcHandler) CreateUser(ctx context.Context, in *userPb.CreateUserRequest) (*userPb.Response, error) {
	log.Printf("Create user received request! %v", in)

	userDto, err := mapUserTypeFromPb(in.GetUser())
	if err != nil {
		return &userPb.Response{}, err
	}

	createdUser, err := s.service.UserStore.CreateOne(ctx, userDto)
	if err != nil {
		return &userPb.Response{}, err
	}

	res, err := mapResponseFromType("User created successfully", createdUser.InsertedID, userDto)
	if err != nil {
		return &userPb.Response{}, err
	}

	return &res, nil
}
