package user

import (
	"context"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
)

func (h *Handler) GetUser(ctx context.Context, in *userPb.UserRequest) (*userPb.UserResponse, error) {
	// translate request
	// validate request

	//consume service
	res, err := h.userService.Get(ctx, in.GetUserId())
	if err != nil {
		return &userPb.UserResponse{}, err
	}

	// translate response
	return &userPb.UserResponse{
		Message: "User found",
		User:    res,
	}, nil
}
