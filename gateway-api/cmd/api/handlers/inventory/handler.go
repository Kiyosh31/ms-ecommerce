package inventory

import (
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/inventory-service"
	"go.uber.org/zap"
)

type Handler struct {
	inventoryServiceClient inventoryPb.InventoryServiceClient
	logger                 *zap.SugaredLogger
}

func NewInventoryHandler(
	inventoryServiceClient inventoryPb.InventoryServiceClient,
	logger *zap.SugaredLogger,
) *Handler {
	return &Handler{
		inventoryServiceClient: inventoryServiceClient,
		logger:                 logger,
	}
}
