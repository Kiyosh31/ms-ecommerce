package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
)

func (s Service) Reactivate(ctx context.Context, email, password string) (*userPb.User, error) {
	// check is user exist
	userToReactivate, exist, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return &userPb.User{}, err
	}
	if !exist {
		return &userPb.User{}, errors.New("user not found")
	}
	if userToReactivate.IsActive {
		return &userPb.User{}, errors.New("user already active")
	}
	err = utils.IsPasswordValid(userToReactivate.Password, password)
	if err != nil {
		// log.Errorf("error incorrect password: %v", err)
		return nil, fmt.Errorf("incorrect password: %v", err)
	}
	// activating user
	userToReactivate.IsActive = true

	// save to DB
	_, err = s.repository.Update(ctx, userToReactivate)
	if err != nil {
		return &userPb.User{}, err
	}

	return &userPb.User{
		Id:        userToReactivate.ID.Hex(),
		Name:      userToReactivate.Name,
		Email:     userToReactivate.Email,
		Password:  userToReactivate.Password,
		Addresses: userutils.MapAddressToProto(userToReactivate.Addresses),
		// Orders: ,
	}, nil
}
