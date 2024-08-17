package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) Create(ctx context.Context, newUser domain.UserSchema) (*userPb.User, error) {
	s.logger.Infof("create user: %v", newUser)

	// check is user exist
	_, exist, err := s.repository.GetByEmail(ctx, newUser.Email)
	if err != nil {
		s.logger.Error(err)
		return &userPb.User{}, err
	}
	if exist {
		err := errors.New("user already exists")
		s.logger.Error(err)
		return &userPb.User{}, err
	}

	res, err := s.repository.Create(ctx, newUser)
	if err != nil {
		s.logger.Error(err)
		return &userPb.User{}, err
	}

	newUserId, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return &userPb.User{}, errors.New("error getting new user id")
	}
	newUser.ID = newUserId

	s.logger.Infof("create user finished with: %v", newUser)
	return &userPb.User{
		Id:        newUser.ID.Hex(),
		Name:      newUser.Name,
		Email:     newUser.Email,
		Password:  newUser.Password,
		Addresses: userutils.MapAddressToProto(newUser.Addresses),
		// Orders: ,
	}, nil
}
