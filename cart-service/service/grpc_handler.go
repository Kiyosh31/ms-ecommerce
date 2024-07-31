package service

import (
	"context"
	"log"

	cartPb "github.com/Kiyosh31/ms-ecommerce/cart-service/proto"
)

func (s *CartService) CreateCart(ctx context.Context, in *cartPb.CreateCartRequest) (*cartPb.SingleCartResponse, error) {
	log.Println("jajatl")
	return &cartPb.SingleCartResponse{}, nil
}

func (s *CartService) GetCart(ctx context.Context, in *cartPb.GetCartRequest) (*cartPb.SingleCartResponse, error) {
	return &cartPb.SingleCartResponse{}, nil
}

func (s *CartService) GetCarts(ctx context.Context, in *cartPb.GetCartsRequest) (*cartPb.MultipleCartResponse, error) {
	return &cartPb.MultipleCartResponse{}, nil
}
