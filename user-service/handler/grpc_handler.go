package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/user-service/service"
	"github.com/Kiyosh31/ms-ecommerce/user-service/user_types"
	"go.mongodb.org/mongo-driver/mongo"

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

func (h *UserServiceGrpcHandler) CreateUser(ctx context.Context, in *userPb.CreateUserRequest) (*userPb.Response, error) {
	log.Printf("Create user received request! %v", in)

	in.GetUser().IsActive = true
	userDto, err := createUserSchemaDto(in.GetUser())
	if err != nil {
		return &userPb.Response{}, err
	}

	foundedUser, err := h.service.UserStore.GetOneByEmail(ctx, in.GetUser().GetEmail())
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return &userPb.Response{}, err
		}
	}
	if !reflect.DeepEqual(foundedUser, user_types.UserSchema{}) {
		return &userPb.Response{}, errors.New("user already exists")
	}

	createdUser, err := h.service.UserStore.CreateOne(ctx, userDto)
	if err != nil {
		return &userPb.Response{}, err
	}

	res, err := createResponsePbDto("User created successfully", createdUser.InsertedID, userDto)
	if err != nil {
		return &userPb.Response{}, err
	}

	return &res, nil
}

func (h *UserServiceGrpcHandler) GetUser(ctx context.Context, in *userPb.GetUserRequest) (*userPb.Response, error) {
	log.Printf("Get user received request! %v", in)

	userID, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		return &userPb.Response{}, err
	}

	foundedUser, err := h.service.UserStore.GetOne(ctx, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &userPb.Response{}, errors.New(utils.NOT_FOUND)
		}
		return &userPb.Response{}, err
	}

	res, err := createResponsePbDto("User found", foundedUser.ID, foundedUser)
	if err != nil {
		return &userPb.Response{}, err
	}

	return &res, nil
}

func (h *UserServiceGrpcHandler) UpdateUser(ctx context.Context, in *userPb.UpdateUserRequest) (*userPb.Response, error) {
	log.Printf("Update user received request! %v", in)

	userID, err := database.GetMongoId(in.GetUser().GetId())
	if err != nil {
		return &userPb.Response{}, err
	}

	foundedUser, err := h.service.UserStore.GetOne(ctx, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &userPb.Response{}, errors.New(utils.NOT_FOUND)
		}
		return &userPb.Response{}, err
	}
	if reflect.DeepEqual(foundedUser, user_types.UserSchema{}) {
		return &userPb.Response{}, errors.New(utils.NOT_FOUND)
	}

	userToUpdate, err := createUserSchemaDto(in.GetUser())
	if err != nil {
		return &userPb.Response{}, err
	}

	_, err = h.service.UserStore.UpdateOne(ctx, userToUpdate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &userPb.Response{}, errors.New(utils.NOT_FOUND)
		}
		return &userPb.Response{}, err
	}

	res, err := createResponsePbDto("User updated successfully", nil, userToUpdate)
	if err != nil {
		return &userPb.Response{}, err
	}

	return &res, nil
}

func (h *UserServiceGrpcHandler) DeleteUser(ctx context.Context, in *userPb.DeleteUserRequest) (*userPb.Response, error) {
	log.Printf("Update user received request! %v", in)

	userID, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		return &userPb.Response{}, err
	}

	foundedUser, err := h.service.UserStore.GetOne(ctx, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &userPb.Response{}, errors.New(utils.NOT_FOUND)
		}
		return &userPb.Response{}, err
	}
	if reflect.DeepEqual(foundedUser, user_types.UserSchema{}) {
		return &userPb.Response{}, errors.New(utils.NOT_FOUND)
	}
	foundedUser.IsActive = false

	_, err = h.service.UserStore.UpdateOne(ctx, foundedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &userPb.Response{}, errors.New(utils.NOT_FOUND)
		}
		return &userPb.Response{}, err
	}

	res, err := createResponsePbDto("User deleted successfully", nil, foundedUser)
	if err != nil {
		return &userPb.Response{}, err
	}

	return &res, nil
}

func (h *UserServiceGrpcHandler) ReactivateUser(ctx context.Context, in *userPb.ReactivarUserRequest) (*userPb.Response, error) {
	log.Printf("reactivate user received request! %v", in)

	foundedUser, err := h.service.UserStore.GetOneDeactivated(ctx, in.GetEmail())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &userPb.Response{}, errors.New(utils.NOT_FOUND)
		}
		return &userPb.Response{}, err
	}
	if reflect.DeepEqual(foundedUser, user_types.UserSchema{}) {
		return &userPb.Response{}, errors.New(utils.NOT_FOUND)
	}

	err = utils.CheckPassword(foundedUser.Password, in.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("incorrect password: %v", err)
	}
	foundedUser.IsActive = true

	_, err = h.service.UserStore.UpdateOne(ctx, foundedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &userPb.Response{}, errors.New(utils.NOT_FOUND)
		}
		return &userPb.Response{}, err
	}

	res, err := createResponsePbDto("User reactivated successfully", nil, foundedUser)
	if err != nil {
		return &userPb.Response{}, err
	}

	return &res, nil
}
