package user

import (
	"context"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
)

func (h *Handler) ReactivateUser(ctx context.Context, in *userPb.ReactivateUserRequest) (*userPb.UserResponse, error) {
	//consume service
	res, err := h.userService.Reactivate(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		return &userPb.UserResponse{}, err
	}

	// translate response
	return &userPb.UserResponse{
		Message: "user reactivated successfully",
		User:    res,
	}, nil
}
