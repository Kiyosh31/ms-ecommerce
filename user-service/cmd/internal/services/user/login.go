package user

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
)

func (s Service) Login(ctx context.Context, email, password string) (*userPb.TokenResponse, error) {
	// check is user exist
	user, exist, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return &userPb.TokenResponse{}, err
	}
	if !exist {
		return &userPb.TokenResponse{}, errors.New("user not found")
	}
	if !user.IsActive {
		return &userPb.TokenResponse{}, errors.New("user deactivated")
	}

	err = utils.IsPasswordValid(user.Password, password)
	if err != nil {
		s.logger.Errorf("error incorrect password: %v", err)
		return nil, fmt.Errorf("incorrect password: %v", err)
	}

	duration, err := strconv.Atoi(s.tokenDurationTime)
	if err != nil {
		s.logger.Errorf("error converting tokenDuration: %v", err)
		return nil, fmt.Errorf("invalid token duration: %v", err)
	}

	accessToken, err := s.tokenCreator.CreateToken(
		user.Email,
		string(user.Role),
		duration,
	)
	if err != nil {
		return &userPb.TokenResponse{}, err
	}

	return &userPb.TokenResponse{
		Token: accessToken,
	}, nil
}
