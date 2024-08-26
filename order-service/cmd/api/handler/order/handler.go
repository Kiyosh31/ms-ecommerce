package order

import (
	"github.com/Kiyosh31/ms-ecommerce/order-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/order-service/cmd/internal/services/order"
	orderPb "github.com/Kiyosh31/ms-ecommerce/order-service/cmd/proto"

	"go.uber.org/zap"
)

type Handler struct {
	orderPb.UnimplementedOrderServiceServer
	orderService ports.OrderService
	logger       *zap.SugaredLogger
}

func NewOrderHandler(orderService *order.Service, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		orderService: orderService,
		logger:       logger,
	}
}
