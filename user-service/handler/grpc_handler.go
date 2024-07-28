package handler

import (
	"context"
	"fmt"
	"log"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/user-service/service"
	"github.com/Kiyosh31/ms-ecommerce/user-service/types"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

	foundedUser, err := s.service.UserStore.GetOneByEmail(ctx, in.GetUser().GetEmail())
	if err != nil {
		return &userPb.Response{}, err
	}
	if foundedUser.Email != "" {
		return &userPb.Response{}, fmt.Errorf("user already exists")
	}
	log.Println(foundedUser)

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

func (s *UserServiceGrpcHandler) GetUser(ctx context.Context, in *userPb.GetUserRequest) (*userPb.Response, error) {
	log.Printf("Get user received request! %v", in)

	userID, err := primitive.ObjectIDFromHex(in.GetUserId())
	if err != nil {
		return &userPb.Response{}, err
	}

	foundedUser, err := s.service.UserStore.GetOne(ctx, userID)
	if err != nil {
		return &userPb.Response{}, err
	}

	var userDto types.UserSchema
	res, err := mapResponseFromType("User founded successfully", foundedUser.ID, userDto)
	if err != nil {
		return &userPb.Response{}, err
	}

	return &res, nil
}
