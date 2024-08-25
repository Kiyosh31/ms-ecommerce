package user

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

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

	// userId, err := strconv.ParseInt(user.ID.Hex(), 16, 64) // Base 16 for hex, 64-bit integer
	// if err != nil {
	// 	s.logger.Errorf("error parsing userId to int: %v", err)
	// 	return nil, err
	// }

	tokenDuration, err := strconv.Atoi(s.tokenDurationTime)
	if err != nil {
		s.logger.Errorf("error converting tokenDuration: %v", err)
		return nil, fmt.Errorf("invalid token duration: %v", err)
	}

	duration := time.Duration(tokenDuration) * time.Minute

	accessToken, _, err := s.tokenCreator.CreateToken(
		user.ID.Hex(),
		user.Email,
		user.IsActive,
		duration,
	)
	if err != nil {
		return &userPb.TokenResponse{}, err
	}

	return &userPb.TokenResponse{
		Token: accessToken,
	}, nil
}
