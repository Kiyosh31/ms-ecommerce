package service

import (
	"context"

	orderPb "github.com/Kiyosh31/ms-ecommerce/order-service/proto"
)

func (s *OrderService) CreateOrder(ctx context.Context, in *orderPb.OrderRequest) (*orderPb.OrderResponse, error) {
	s.logger.Infof("create order request incoming: %v", in.GetOrder())
	return &orderPb.OrderResponse{}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, in *orderPb.OrderRequest) (*orderPb.OrderResponse, error) {
	return &orderPb.OrderResponse{}, nil

}

func (s *OrderService) CancelOrder(ctx context.Context, in *orderPb.OrderRequest) (*orderPb.OrderResponse, error) {
	return &orderPb.OrderResponse{}, nil

}
