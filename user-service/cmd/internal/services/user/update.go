package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
)

func (s Service) Update(ctx context.Context, user domain.UserSchema) (*userPb.User, error) {
	// validate user exists
	existingUser, exist, err := s.repository.Get(ctx, user.ID)
	if err != nil {
		return &userPb.User{}, err
	}
	if !exist {
		return &userPb.User{}, errors.New("user not found")
	}
	if !existingUser.IsActive {
		return &userPb.User{}, errors.New("user is deactivated")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return &userPb.User{}, err
	}
	user.Password = hashedPassword

	//update to db
	_, err = s.repository.Update(ctx, user)
	if err != nil {
		return &userPb.User{}, err
	}

	// translate response
	res := &userPb.User{
		Id:        user.ID.Hex(),
		Name:      user.Name,
		Password:  user.Password,
		Email:     user.Email,
		Addresses: userutils.MapAddressToProto(user.Addresses),
	}

	return res, nil
}
