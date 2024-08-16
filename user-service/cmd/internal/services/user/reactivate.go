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
	user_to_reactivate, exist, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return &userPb.User{}, err
	}
	if !exist {
		return &userPb.User{}, errors.New("user not found")
	}
	if user_to_reactivate.IsActive {
		return &userPb.User{}, errors.New("user already active")
	}
	err = utils.CheckPassword(user_to_reactivate.Password, password)
	if err != nil {
		// log.Errorf("error incorrect password: %v", err)
		return nil, fmt.Errorf("incorrect password: %v", err)
	}
	// activating user
	user_to_reactivate.IsActive = true

	// save to DB
	_, err = s.repository.Update(ctx, user_to_reactivate)
	if err != nil {
		return &userPb.User{}, err
	}

	return &userPb.User{
		Id:        user_to_reactivate.ID.Hex(),
		Name:      user_to_reactivate.Name,
		Email:     user_to_reactivate.Email,
		Password:  user_to_reactivate.Password,
		Addresses: userutils.MapAddressToProto(user_to_reactivate.Addresses),
		// Orders: ,
	}, nil
}
