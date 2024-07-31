package service

import (
	"context"
	"errors"
	"log"
	"reflect"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/cart-service/cart_types"
	cartPb "github.com/Kiyosh31/ms-ecommerce/cart-service/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *CartService) CreateCart(ctx context.Context, in *cartPb.CreateCartRequest) (*cartPb.SingleCartResponse, error) {
	log.Printf("Create cart request received! %v", in)

	cartDto, err := createCartSchemaDto(in.GetCart())
	if err != nil {
		return &cartPb.SingleCartResponse{}, err
	}

	createdCart, err := s.CartStore.CreateOne(ctx, cartDto)
	if err != nil {
		return &cartPb.SingleCartResponse{}, err
	}

	res, err := createSingleCartResponseDto("cart created successfully", createdCart.InsertedID, cartDto)
	if err != nil {
		return &cartPb.SingleCartResponse{}, err
	}

	return res, nil
}

func (s *CartService) GetCart(ctx context.Context, in *cartPb.GetCartRequest) (*cartPb.SingleCartResponse, error) {
	log.Printf("Get cart request received! %v", in)

	userId, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		return &cartPb.SingleCartResponse{}, err
	}

	cartId, err := database.GetMongoId(in.GetCartId())
	if err != nil {
		return &cartPb.SingleCartResponse{}, err
	}

	foundedCart, err := s.CartStore.GetOne(ctx, cartId)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return &cartPb.SingleCartResponse{}, err
		}
		return &cartPb.SingleCartResponse{}, err
	}
	if reflect.DeepEqual(foundedCart, cart_types.CartSchema{}) {
		return &cartPb.SingleCartResponse{}, errors.New("cart not found")
	}
	if userId != foundedCart.UserId {
		return &cartPb.SingleCartResponse{}, errors.New("cart does not belong to this user")
	}

	res, err := createSingleCartResponseDto("Cart founded", nil, foundedCart)
	if err != nil {
		return &cartPb.SingleCartResponse{}, err
	}

	return res, nil
}

func (s *CartService) GetAllCarts(ctx context.Context, in *cartPb.GetCartsRequest) (*cartPb.MultipleCartResponse, error) {
	log.Printf("Get all carts request received! %v", in)

	userId, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		return &cartPb.MultipleCartResponse{}, err
	}

	foundedCarts, err := s.CartStore.GetAll(ctx, userId)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return &cartPb.MultipleCartResponse{}, err
		}
		return &cartPb.MultipleCartResponse{}, err
	}

	res := createMultipleCartResponseDto("Cart founded", foundedCarts)

	return res, nil
}
