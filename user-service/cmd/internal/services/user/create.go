package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) Create(ctx context.Context, new_user domain.UserSchema) (*userPb.User, error) {
	s.logger.Infof("create user: %v", new_user)

	// check is user exist
	_, exist, err := s.repository.GetByEmail(ctx, new_user.Email)
	if err != nil {
		s.logger.Error(err)
		return &userPb.User{}, err
	}
	if exist {
		err := errors.New("user already exists")
		s.logger.Error(err)
		return &userPb.User{}, err
	}

	res, err := s.repository.Create(ctx, new_user)
	if err != nil {
		s.logger.Error(err)
		return &userPb.User{}, err
	}

	new_user_id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return &userPb.User{}, errors.New("error getting new user id")
	}
	new_user.ID = new_user_id

	s.logger.Infof("create user finished with: %v", new_user)
	return &userPb.User{
		Id:        new_user.ID.Hex(),
		Name:      new_user.Name,
		Email:     new_user.Email,
		Password:  new_user.Password,
		Addresses: userutils.MapAddressToProto(new_user.Addresses),
		// Orders: ,
	}, nil
}
