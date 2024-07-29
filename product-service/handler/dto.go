package handler

import (
	"github.com/Kiyosh31/ms-ecommerce-common/database"

	"github.com/Kiyosh31/ms-ecommerce/product-service/product_types"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createProductSchemaDto(in *productPb.Product) (product_types.ProductSchema, error) {
	var productId primitive.ObjectID
	var err error

	if in.GetId() != "" {
		productId, err = database.GetMongoId(in.GetId())
		if err != nil {
			return product_types.ProductSchema{}, err
		}
	}

	sellerId, err := database.GetMongoId(in.GetSellerId())
	if err != nil {
		return product_types.ProductSchema{}, err
	}

	return product_types.ProductSchema{
		ID:          productId,
		Name:        in.GetName(),
		Price:       in.GetPrice(),
		Description: in.GetDescription(),
		SellerID:    sellerId,
	}, nil
}
