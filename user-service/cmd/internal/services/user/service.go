package user

import (
	"github.com/Kiyosh31/ms-ecommerce-common/token"
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/repositories/mongo/user"
	"go.uber.org/zap"
)

type Service struct {
	repository        ports.UserRepository
	logger            *zap.SugaredLogger
	tokenCreator      *token.JwtCreator
	tokenDurationTime string
}

func NewUserService(
	repository *user.Repository,
	logger *zap.SugaredLogger,
	secretKey string,
	tokenDurationTime string,
) *Service {
	return &Service{
		repository:        repository,
		logger:            logger,
		tokenCreator:      token.NewJwtCreator(secretKey),
		tokenDurationTime: tokenDurationTime,
	}
}
