package payment

import (
	"github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/internal/repositories/mongo/payment"
	"go.uber.org/zap"
)

type Service struct {
	repository ports.PaymentRepository
	logger     *zap.SugaredLogger
}

func NewUserService(
	repository *payment.Repository,
	logger *zap.SugaredLogger,
) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}
