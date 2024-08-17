package user

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
)

func (h *Handler) CreateUser(ctx context.Context, in *userPb.UserRequest) (*userPb.UserResponse, error) {
	h.logger.Infof("create user request incoming: %v", in)
	// translate request
	// validate request
	hashedPassword, err := utils.HashPassword(in.GetUser().GetPassword())
	if err != nil {
		return &userPb.UserResponse{}, err
	}

	addresses, err := userutils.MapAddressToDomain(in.GetUser().GetAddresses())
	if err != nil {
		return &userPb.UserResponse{}, nil
	}

	newUser := domain.UserSchema{
		Name:      in.GetUser().GetName(),
		Email:     in.GetUser().GetEmail(),
		Password:  hashedPassword,
		Addresses: addresses,
		// Orders: ,
		IsActive: true,
	}

	//consume service
	res, err := h.userService.Create(ctx, newUser)
	if err != nil {
		return &userPb.UserResponse{}, err
	}

	// translate response
	return &userPb.UserResponse{
		Message: "user created successfully",
		User:    res,
	}, nil
}
