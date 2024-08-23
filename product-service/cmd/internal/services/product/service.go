package product

import (
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/ports"
	"go.uber.org/zap"
)

type Service struct {
	repository ports.ProductRepository
	logger     *zap.SugaredLogger
}

func NewProductService(
	repository ports.ProductRepository,
	logger *zap.SugaredLogger,
) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}
