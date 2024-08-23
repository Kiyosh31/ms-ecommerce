package category

import (
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/ports"
	"go.uber.org/zap"
)

type Service struct {
	repository ports.CategoryRepository
	logger     *zap.SugaredLogger
}

func NewCategoryService(
	repository ports.CategoryRepository,
	logger *zap.SugaredLogger,
) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}
