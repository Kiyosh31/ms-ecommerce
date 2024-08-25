package user

import (
	"github.com/Kiyosh31/ms-ecommerce-common/token"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
	"go.uber.org/zap"
)

type Handler struct {
	userServiceClient userPb.UserServiceClient
	logger            *zap.SugaredLogger
	tokenCreator      *token.JwtCreator
}

func NewUserHandler(
	userServiceClient userPb.UserServiceClient,
	logger *zap.SugaredLogger,
	secretKey string,
) *Handler {
	return &Handler{
		userServiceClient: userServiceClient,
		logger:            logger,
		tokenCreator:      token.NewJwtCreator(secretKey),
	}
}
