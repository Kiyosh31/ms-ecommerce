package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
)

func (s *Service) Get(ctx context.Context, id string) (*userPb.User, error) {
	userId, err := database.GetMongoId(id)
	if err != nil {
		return &userPb.User{}, err
	}

	user, exist, err := s.repository.Get(ctx, userId)
	if err != nil {
		return &userPb.User{}, err
	}
	if !exist {
		return &userPb.User{}, errors.New("user not found")
	}
	if !user.IsActive {
		return &userPb.User{}, errors.New("user is deactivated")
	}

	res := &userPb.User{
		Id:        user.ID.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Addresses: userutils.MapAddressToProto(user.Addresses),
		// Orders:    user.Orders,
	}

	return res, nil
}
