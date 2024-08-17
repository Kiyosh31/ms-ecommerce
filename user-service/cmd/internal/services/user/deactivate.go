package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
)

func (s Service) Deactivate(ctx context.Context, id string) (bool, error) {
	userId, err := database.GetMongoId(id)
	if err != nil {
		return false, err
	}

	// check is user exist
	userToDeactivate, exist, err := s.repository.Get(ctx, userId)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, errors.New("user not found")
	}
	if !userToDeactivate.IsActive {
		return false, errors.New("user already deactivated")
	}
	// deactivating user
	userToDeactivate.IsActive = false

	// save to DB
	_, err = s.repository.Update(ctx, userToDeactivate)
	if err != nil {
		return false, err
	}

	return true, nil
}
