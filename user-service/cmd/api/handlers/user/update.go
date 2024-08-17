package user

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	userutils "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/user_utils"
)

func (h *Handler) UpdateUser(ctx context.Context, in *userPb.UserRequest) (*userPb.UserResponse, error) {
	// translate request
	// validate request
	userId, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		return &userPb.UserResponse{}, err
	}

	addresses, err := userutils.MapAddressToDomain(in.GetUser().GetAddresses())
	if err != nil {
		return &userPb.UserResponse{}, err
	}

	userToUpdate := domain.UserSchema{
		ID:        userId,
		Name:      in.GetUser().GetName(),
		Email:     in.GetUser().GetEmail(),
		Password:  in.GetUser().GetEmail(),
		Addresses: addresses,
		IsActive:  true,
	}

	//consume service
	res, err := h.userService.Update(ctx, userToUpdate)
	if err != nil {
		return &userPb.UserResponse{}, err
	}

	// translate response
	return &userPb.UserResponse{
		Message: "user updated successfully",
		User:    res,
	}, nil
}
