package service

import (
	orderPb "github.com/Kiyosh31/ms-ecommerce/order-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/order-service/store"
	"go.uber.org/zap"
)

type OrderService struct {
	orderPb.UnimplementedOrderServiceServer
	GrpcAdd      string
	ProductStore store.OrderStore
	logger       *zap.SugaredLogger
}

func NewOrderService(
	GrpcAdd string,
	ProductStore store.OrderStore,
	logger *zap.SugaredLogger,
) *OrderService {
	return &OrderService{
		GrpcAdd:      GrpcAdd,
		ProductStore: ProductStore,
		logger:       logger,
	}
}
