package service

import (
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/inventory-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/store"
	"go.uber.org/zap"
)

type InventoryService struct {
	inventoryPb.UnimplementedInventoryServiceServer
	GrpcAddr       string
	InventoryStore store.InventoryStore
	logger         *zap.SugaredLogger
}

func NewInventoryService(
	GrpcAddr string,
	InventoryStore store.InventoryStore,
	logger *zap.SugaredLogger,
) *InventoryService {
	return &InventoryService{
		GrpcAddr:       GrpcAddr,
		InventoryStore: InventoryStore,
		logger:         logger,
	}
}
