package inventory

import "go.uber.org/zap"

type Service struct {
	inventoryClientGrpcAddr string
	logger                  *zap.SugaredLogger
}

func NewInventoryService(inventoryClientGrpcAddr string,
	logger *zap.SugaredLogger) *Service {
	return &Service{
		inventoryClientGrpcAddr: inventoryClientGrpcAddr,
		logger:                  logger,
	}
}
