package service

import (
	"context"
	"errors"
	"reflect"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/cart-service/cart_types"
	cartPb "github.com/Kiyosh31/ms-ecommerce/cart-service/proto"
)

func (s *CartService) CreateCart(ctx context.Context, in *cartPb.CreateCartRequest) (*cartPb.SingleCartResponse, error) {
	s.logger.Infof("create cart incoming request: %v", in)

	cartDto, err := createCartSchemaDto(in.GetCart())
	if err != nil {
		s.logger.Errorf("error creating cart dto: %v", in)
		return &cartPb.SingleCartResponse{}, err
	}

	createdCart, err := s.CartStore.CreateOne(ctx, cartDto)
	if err != nil {
		s.logger.Errorf("error creating cart: %v", err)
		return &cartPb.SingleCartResponse{}, err
	}

	res, err := createSingleCartResponseDto("cart created successfully", createdCart.InsertedID, cartDto)
	if err != nil {
		s.logger.Errorf("error creating cart response: %v", err)
		return &cartPb.SingleCartResponse{}, err
	}

	s.logger.Infof("create cart request finished: %v", res)
	return res, nil
}

func (s *CartService) GetCart(ctx context.Context, in *cartPb.GetCartRequest) (*cartPb.SingleCartResponse, error) {
	s.logger.Infof("Get cart incoming request: %v", in)

	userId, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		s.logger.Errorf("error getting cartId: %v", err)
		return &cartPb.SingleCartResponse{}, err
	}

	cartId, err := database.GetMongoId(in.GetCartId())
	if err != nil {
		s.logger.Errorf("error getting cartId: %v", err)
		return &cartPb.SingleCartResponse{}, err
	}

	foundedCart, err := s.CartStore.GetOne(ctx, cartId)
	if err != nil {
		s.logger.Errorf("error getting cart: %v", err)
		return &cartPb.SingleCartResponse{}, err
	}
	if reflect.DeepEqual(foundedCart, cart_types.CartSchema{}) {
		s.logger.Errorf("cart not found: %v", err)
		return &cartPb.SingleCartResponse{}, errors.New("cart not found")
	}
	if userId != foundedCart.UserId {
		s.logger.Errorf("cart does not belong to this user: %v", err)
		return &cartPb.SingleCartResponse{}, errors.New("cart does not belong to this user")
	}

	res, err := createSingleCartResponseDto("Cart founded", nil, foundedCart)
	if err != nil {
		s.logger.Errorf("error creating cart response: %v", err)
		return &cartPb.SingleCartResponse{}, err
	}

	s.logger.Infof("get cart request finished: %v", err)
	return res, nil
}

func (s *CartService) GetAllCarts(ctx context.Context, in *cartPb.GetCartsRequest) (*cartPb.MultipleCartResponse, error) {
	s.logger.Infof("get all carts incoming request: %v", in)

	userId, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		s.logger.Errorf("error getting cartId: %v", err)
		return &cartPb.MultipleCartResponse{}, err
	}

	foundedCarts, err := s.CartStore.GetAll(ctx, userId)
	if err != nil {
		s.logger.Errorf("error getting all cart: %v", err)
		return &cartPb.MultipleCartResponse{}, err
	}

	res := createMultipleCartResponseDto("Cart founded", foundedCarts)

	s.logger.Infof("get all cart request finished: %v", err)
	return res, nil
}
