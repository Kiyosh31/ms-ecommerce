package user

import (
	"context"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
)

func (h *Handler) DeactivateUser(ctx context.Context, in *userPb.UserRequest) (*userPb.UserResponse, error) {
	// translate request
	// validate request

	//consume service
	res, err := h.userService.Deactivate(ctx, in.GetUserId())
	if err != nil {
		return &userPb.UserResponse{}, err
	}

	var message string
	if res {
		message = "user deactivated successfully"
	} else {
		message = "problem deactivating user"
	}

	// translate response
	return &userPb.UserResponse{
		Message: message,
	}, nil
}
