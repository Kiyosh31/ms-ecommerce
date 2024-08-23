package brand

import (
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/ports"
	"go.uber.org/zap"
)

type Service struct {
	repository ports.BrandRepository
	logger     *zap.SugaredLogger
}

func NewBrandService(
	repository ports.BrandRepository,
	logger *zap.SugaredLogger,
) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}
