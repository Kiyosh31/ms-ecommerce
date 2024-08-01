package service

import (
	cartPb "github.com/Kiyosh31/ms-ecommerce/cart-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/cart-service/store"
	"go.uber.org/zap"
)

type CartService struct {
	cartPb.UnimplementedCartServiceServer
	GrpcAddr  string
	CartStore store.CartStore
	logger    *zap.SugaredLogger
}

func NewCartService(grpcAddr string, cartStore store.CartStore, logger *zap.SugaredLogger) *CartService {
	return &CartService{
		GrpcAddr:  grpcAddr,
		CartStore: cartStore,
		logger:    logger,
	}
}
