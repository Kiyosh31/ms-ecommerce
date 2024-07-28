package handler

import userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"

func mapCreateUserRequestToPb(in *userPb.User) *userPb.CreateUserRequest {
	return &userPb.CreateUserRequest{
		User: in,
	}
}
