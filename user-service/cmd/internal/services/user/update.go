package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
)

func (s Service) Update(ctx context.Context, user_to_update domain.UserSchema) (*userPb.User, error) {
	// validate user exists
	existing_user, exist, err := s.repository.Get(ctx, user_to_update.ID)
	if err != nil {
		return &userPb.User{}, err
	}
	if !exist {
		return &userPb.User{}, errors.New("user not found")
	}

	//update to db
	_, err = s.repository.Update(ctx, user_to_update)
	if err != nil {
		return &userPb.User{}, err
	}

	// translate response
	res := &userPb.User{
		Id:        existing_user.ID.Hex(),
		Name:      existing_user.Name,
		Password:  existing_user.Password,
		Email:     existing_user.Email,
		Addresses: userutils.MapAddressToProto(existing_user.Addresses),
	}

	return res, nil
}
