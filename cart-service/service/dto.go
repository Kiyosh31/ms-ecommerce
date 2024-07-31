package service

import (
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/cart-service/cart_types"
	cartPb "github.com/Kiyosh31/ms-ecommerce/cart-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createCartSchemaDto(in *cartPb.Cart) (cart_types.CartSchema, error) {
	var cartId primitive.ObjectID
	var err error

	if in.Id != "" {
		cartId, err = database.GetMongoId(in.GetId())
		if err != nil {
			return cart_types.CartSchema{}, err
		}
	}

	userId, err := database.GetMongoId(in.GetUserId())
	if err != nil {
		return cart_types.CartSchema{}, err
	}

	return cart_types.CartSchema{
		ID:     cartId,
		UserId: userId,
		Total:  float32(in.GetTotal()),
	}, nil
}

func createSingleCartResponseDto(message string, id interface{}, in cart_types.CartSchema) (*cartPb.SingleCartResponse, error) {
	var cartId primitive.ObjectID
	var ok bool

	if id != nil {
		cartId, ok = id.(primitive.ObjectID)
		if !ok {
			return &cartPb.SingleCartResponse{}, fmt.Errorf("failed to parse _id to string")
		}
	} else {
		cartId = in.ID
	}

	return &cartPb.SingleCartResponse{
		Message: message,
		Cart: &cartPb.Cart{
			Id:     cartId.Hex(),
			UserId: in.ID.Hex(),
			Total:  float64(in.Total),
		},
	}, nil
}

func createMultipleCartResponseDto(message string, in []cart_types.CartSchema) *cartPb.MultipleCartResponse {
	var carts []*cartPb.Cart

	for _, cart := range in {
		c := cartPb.Cart{
			Id:     cart.ID.Hex(),
			UserId: cart.UserId.Hex(),
			Total:  float64(cart.Total),
		}

		carts = append(carts, &c)
	}

	return &cartPb.MultipleCartResponse{
		Message: message,
		Cart:    carts,
	}
}
