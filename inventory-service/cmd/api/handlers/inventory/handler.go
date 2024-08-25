package inventory

import (
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/services/inventory"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/proto"
	"go.uber.org/zap"
)

type Handler struct {
	inventoryPb.UnimplementedInventoryServiceServer
	userService ports.InventoryService
	logger      *zap.SugaredLogger
}

func NewInventoryHandler(userService *inventory.Service, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		userService: userService,
		logger:      logger,
	}
}
