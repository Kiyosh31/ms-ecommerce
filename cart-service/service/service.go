package service

import (
	cartPb "github.com/Kiyosh31/ms-ecommerce/cart-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/cart-service/store"
)

type CartService struct {
	cartPb.UnimplementedCartServiceServer
	GrpcAddr  string
	CartStore store.CartStore
}

func NewCartService(grpcAddr string, cartStore store.CartStore) *CartService {
	return &CartService{
		GrpcAddr:  grpcAddr,
		CartStore: cartStore,
	}
}
