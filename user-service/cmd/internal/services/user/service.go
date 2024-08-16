package user

import (
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/repositories/user"
	"go.uber.org/zap"
)

type Service struct {
	repository ports.UserRepository
	logger     *zap.SugaredLogger
}

func NewUserService(repository user.Repository, logger *zap.SugaredLogger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}
