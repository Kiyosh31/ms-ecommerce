package handler

import (
	"context"
	"log"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/user-service/service"
	"github.com/Kiyosh31/ms-ecommerce/user-service/user_types"
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
		log.Printf("Parada 1: %v", err)
		return &userPb.Response{}, err
	}

	// foundedUser, err := s.service.UserStore.GetOneByEmail(ctx, in.GetUser().GetEmail())
	// if err == nil && &foundedUser != nil {
	// 	return nil, fmt.Errorf("user already exists")
	// }

	createdUser, err := s.service.UserStore.CreateOne(ctx, userDto)
	log.Printf("Parada 2: %v", createdUser)
	if err != nil {
		log.Printf("Parada 3: %v", err)
		return &userPb.Response{}, err
	}

	res, err := mapResponseFromType("User created successfully", createdUser.InsertedID, userDto)
	log.Printf("Parada 4: %v", &res)
	if err != nil {
		log.Printf("Parada 5: %v", err)
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

	var userDto user_types.UserSchema
	res, err := mapResponseFromType("User founded successfully", foundedUser.ID, userDto)
	if err != nil {
		return &userPb.Response{}, err
	}

	return &res, nil
}
