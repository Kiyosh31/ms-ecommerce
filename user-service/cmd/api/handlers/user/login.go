package user

import (
	"context"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
)

func (h *Handler) LoginUser(ctx context.Context, in *userPb.CredentialsUserRequest) (*userPb.TokenResponse, error) {
	//consume service
	res, err := h.userService.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		return &userPb.TokenResponse{}, err
	}

	// translate response
	return res, nil
}
