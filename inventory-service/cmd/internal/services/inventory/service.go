package inventory

import (
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/repositories/mongo/inventory"
	"go.uber.org/zap"
)

type Service struct {
	repository ports.InventoryRepository
	logger     *zap.SugaredLogger
}

func NewInventoryService(repository *inventory.Repository, logger *zap.SugaredLogger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}
