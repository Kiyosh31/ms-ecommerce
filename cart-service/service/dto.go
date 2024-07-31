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

	var products []cart_types.Product
	for _, prod := range in.GetProducts() {
		p := cart_types.Product{
			ProductId: prod.GetProductId(),
			Quantity:  int64(prod.GetQuantity()),
			Price:     float64(prod.GetPrice()),
		}

		products = append(products, p)
	}

	return cart_types.CartSchema{
		ID:       cartId,
		UserId:   userId,
		Total:    float32(in.GetTotal()),
		Products: products,
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

	var products []*cartPb.Product
	for _, prod := range in.Products {
		p := cartPb.Product{
			ProductId: prod.ProductId,
			Quantity:  float32(prod.Quantity),
			Price:     float32(prod.Price),
		}

		products = append(products, &p)
	}

	return &cartPb.SingleCartResponse{
		Message: message,
		Cart: &cartPb.Cart{
			Id:       cartId.Hex(),
			UserId:   in.ID.Hex(),
			Total:    float64(in.Total),
			Products: products,
		},
	}, nil
}

func createMultipleCartResponseDto(message string, in []cart_types.CartSchema) *cartPb.MultipleCartResponse {
	var carts []*cartPb.Cart
	var products []*cartPb.Product

	for _, cart := range in {
		for _, prod := range cart.Products {
			p := cartPb.Product{
				ProductId: prod.ProductId,
				Quantity:  float32(prod.Quantity),
				Price:     float32(prod.Price),
			}

			products = append(products, &p)
		}

		c := cartPb.Cart{
			Id:       cart.ID.Hex(),
			UserId:   cart.UserId.Hex(),
			Total:    float64(cart.Total),
			Products: products,
		}

		carts = append(carts, &c)
	}

	return &cartPb.MultipleCartResponse{
		Message: message,
		Cart:    carts,
	}
}
