package user

import (
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/services/user"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	"go.uber.org/zap"
)

type Handler struct {
	userPb.UnimplementedUserServiceServer
	userService ports.UserService
	logger      *zap.SugaredLogger
}

func NewUserHandler(userService user.Service, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		userService: userService,
		logger:      logger,
	}
}
