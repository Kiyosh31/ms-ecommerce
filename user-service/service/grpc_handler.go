package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/user-service/user_types"
)

func (s *UserService) CreateUser(ctx context.Context, in *userPb.UserRequest) (*userPb.UserResponse, error) {
	s.logger.Infof("Create user incoming request: %v", in)

	in.GetUser().IsActive = true
	userDto, err := createUserSchemaDto(in.GetUser())
	if err != nil {
		s.logger.Errorf("error unmarshalling body: %v", err)
		return &userPb.UserResponse{}, err
	}

	exists, err := s.UserStore.UserExists(ctx, in.GetUser().GetEmail())
	if err != nil {
		s.logger.Errorf("error finding existing user: %v", err)
		return &userPb.UserResponse{}, err
	}
	if exists {
		s.logger.Errorf("error user exists: %v", err)
		return &userPb.UserResponse{}, errors.New("user already exists")
	}

	createdUser, err := s.UserStore.CreateOne(ctx, userDto)
	if err != nil {
		s.logger.Errorf("error creating user: %v", err)
		return &userPb.UserResponse{}, err
	}

	res, err := createResponsePbDto("User created successfully", createdUser.InsertedID, userDto)
	if err != nil {
		s.logger.Errorf("error creating response: %v", err)
		return &userPb.UserResponse{}, err
	}

	s.logger.Infof("create user request finished: %v", err)
	return &res, nil
}

func (s *UserService) GetUser(ctx context.Context, in *userPb.UserRequest) (*userPb.UserResponse, error) {
	s.logger.Infof("Get user request incoming: %v", in)

	userID, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		s.logger.Errorf("error getting userId: %v", err)
		return &userPb.UserResponse{}, err
	}

	foundedUser, err := s.UserStore.GetOne(ctx, userID)
	if err != nil {
		s.logger.Errorf("error finding user: %v", err)
		return &userPb.UserResponse{}, err
	}

	res, err := createResponsePbDto("User found", foundedUser.ID, foundedUser)
	if err != nil {
		s.logger.Errorf("error creating response: %v", err)
		return &userPb.UserResponse{}, err
	}

	s.logger.Infof("get user request finished: %v", err)
	return &res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, in *userPb.UserRequest) (*userPb.UserResponse, error) {
	s.logger.Infof("Update user incoming request: %v", in)

	exists, err := s.UserStore.UserExists(ctx, in.GetUser().GetEmail())
	if err != nil {
		s.logger.Errorf("error finding existing user: %v", err)
		return &userPb.UserResponse{}, err
	}
	if exists {
		s.logger.Errorf("error user exists: %v", err)
		return &userPb.UserResponse{}, errors.New("user already exists")
	}

	userToUpdate, err := createUserSchemaDto(in.GetUser())
	if err != nil {
		s.logger.Errorf("error creating schema dto: %v", err)
		return &userPb.UserResponse{}, err
	}

	_, err = s.UserStore.UpdateOne(ctx, userToUpdate)
	if err != nil {
		s.logger.Errorf("error updating user: %v", err)
		return &userPb.UserResponse{}, err
	}

	res, err := createResponsePbDto("User updated successfully", nil, userToUpdate)
	if err != nil {
		s.logger.Errorf("error creating response: %v", err)
		return &userPb.UserResponse{}, err
	}

	s.logger.Infof("update user request finished: %v", err)
	return &res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, in *userPb.UserRequest) (*userPb.UserResponse, error) {
	s.logger.Infof("delete user request incoming: %v", in)

	userID, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		s.logger.Errorf("error getting userId: %v", err)
		return &userPb.UserResponse{}, err
	}

	foundedUser, err := s.UserStore.GetOne(ctx, userID)
	if err != nil {
		s.logger.Errorf("error finding user: %v", err)
		return &userPb.UserResponse{}, err
	}
	if reflect.DeepEqual(foundedUser, user_types.UserSchema{}) {
		s.logger.Errorf("existing user: %v", err)
		return &userPb.UserResponse{}, errors.New(utils.NOT_FOUND)
	}
	foundedUser.IsActive = false

	_, err = s.UserStore.UpdateOne(ctx, foundedUser)
	if err != nil {
		s.logger.Errorf("error updating user: %v", err)
		return &userPb.UserResponse{}, err
	}

	res, err := createResponsePbDto("User deleted successfully", nil, foundedUser)
	if err != nil {
		s.logger.Errorf("error creating response: %v", err)
		return &userPb.UserResponse{}, err
	}

	s.logger.Infof("delete user request finished: %v", err)
	return &res, nil
}

func (s *UserService) ReactivateUser(ctx context.Context, in *userPb.ReactivateUserRequest) (*userPb.UserResponse, error) {
	s.logger.Infof("reactivate user incoming request: %v", in)

	foundedUser, err := s.UserStore.GetOneDeactivated(ctx, in.GetEmail())
	if err != nil {
		s.logger.Errorf("error getting deactivated user: %v", err)
		return &userPb.UserResponse{}, err
	}
	if reflect.DeepEqual(foundedUser, user_types.UserSchema{}) {
		s.logger.Errorf("error user active: %v", err)
		return &userPb.UserResponse{}, errors.New(utils.NOT_FOUND)
	}

	err = utils.CheckPassword(foundedUser.Password, in.GetPassword())
	if err != nil {
		s.logger.Errorf("error incorrect password: %v", err)
		return nil, fmt.Errorf("incorrect password: %v", err)
	}
	foundedUser.IsActive = true

	_, err = s.UserStore.UpdateOne(ctx, foundedUser)
	if err != nil {
		s.logger.Errorf("error updating user: %v", err)
		return &userPb.UserResponse{}, err
	}

	res, err := createResponsePbDto("User reactivated successfully", nil, foundedUser)
	if err != nil {
		s.logger.Errorf("error creating response: %v", err)
		return &userPb.UserResponse{}, err
	}

	s.logger.Infof("reactivare user request finished: %v", err)
	return &res, nil
}
