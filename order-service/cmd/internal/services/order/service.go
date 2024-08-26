package order

import (
	"github.com/Kiyosh31/ms-ecommerce/order-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/order-service/cmd/internal/repositories/mongo/order"
	"go.uber.org/zap"
)

type Service struct {
	repository ports.OrderRepository
	logger     *zap.SugaredLogger
}

func NewOrderService(
	repository *order.Repository,
	logger *zap.SugaredLogger,
) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}
