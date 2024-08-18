package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
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
	if !existingUser.IsActive {
		return &userPb.User{}, errors.New("user is deactivated")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(userToUpdate.Password)
	if err != nil {
		return &userPb.User{}, err
	}
	userToUpdate.Password = hashedPassword

	//update to db
	_, err = s.repository.Update(ctx, userToUpdate)
	if err != nil {
		return &userPb.User{}, err
	}

	// translate response
	res := &userPb.User{
		Id:        userToUpdate.ID.Hex(),
		Name:      userToUpdate.Name,
		Password:  userToUpdate.Password,
		Email:     userToUpdate.Email,
		Addresses: userutils.MapAddressToProto(userToUpdate.Addresses),
	}

	return res, nil
}
