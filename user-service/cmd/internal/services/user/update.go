package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
)

func (s Service) Update(ctx context.Context, userToUpdate domain.UserSchema) (*userPb.User, error) {
	// validate user exists
	existingUser, exist, err := s.repository.Get(ctx, userToUpdate.ID)
	if err != nil {
		return &userPb.User{}, err
	}
	if !exist {
		return &userPb.User{}, errors.New("user not found")
	}

	//update to db
	_, err = s.repository.Update(ctx, userToUpdate)
	if err != nil {
		return &userPb.User{}, err
	}

	// translate response
	res := &userPb.User{
		Id:        existingUser.ID.Hex(),
		Name:      existingUser.Name,
		Password:  existingUser.Password,
		Email:     existingUser.Email,
		Addresses: userutils.MapAddressToProto(existingUser.Addresses),
	}

	return res, nil
}
