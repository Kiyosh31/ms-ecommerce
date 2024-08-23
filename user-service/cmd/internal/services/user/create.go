package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) Create(ctx context.Context, user domain.UserSchema) (*userPb.User, error) {
	s.logger.Infof("create user: %v", user)

	// check is user exist
	_, exist, err := s.repository.GetByEmail(ctx, user.Email)
	if err != nil {
		s.logger.Error(err)
		return &userPb.User{}, err
	}
	if exist {
		err := errors.New("user already exists")
		s.logger.Error(err)
		return &userPb.User{}, err
	}

	res, err := s.repository.Create(ctx, user)
	if err != nil {
		s.logger.Error(err)
		return &userPb.User{}, err
	}

	newUserId, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return &userPb.User{}, errors.New("error getting new user id")
	}
	user.ID = newUserId

	s.logger.Infof("create user finished with: %v", user)
	return &userPb.User{
		Id:        user.ID.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Addresses: userutils.MapAddressToProto(user.Addresses),
		// Orders: ,
	}, nil
}
