package user

import (
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
	"go.uber.org/zap"
)

type Handler struct {
	userServiceClient userPb.UserServiceClient
	logger            *zap.SugaredLogger
}

func NewUserHandler(
	userServiceClient userPb.UserServiceClient,
	logger *zap.SugaredLogger,
) *Handler {
	return &Handler{
		userServiceClient: userServiceClient,
		logger:            logger,
	}
}
