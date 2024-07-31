package service

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
		SellerID:    sellerId,
		Name:        in.GetName(),
		Price:       in.GetPrice(),
		Description: in.GetDescription(),
	}, nil
}

func createProductResponseDto(message string, in product_types.ProductSchema) *productPb.ProductResponse {
	return &productPb.ProductResponse{
		Message: message,
		Product: &productPb.Product{
			Id:          in.ID.Hex(),
			Name:        in.Name,
			Price:       in.Price,
			Description: in.Description,
			SellerId:    in.SellerID.Hex(),
		},
	}
}

func createMultipleProductsResponseDto(message string, in []product_types.ProductSchema) *productPb.MultipleProductResponse {
	var products []*productPb.Product

	for _, product := range in {
		prod := productPb.Product{
			Id:          product.ID.Hex(),
			SellerId:    product.SellerID.Hex(),
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
		}

		products = append(products, &prod)
	}

	return &productPb.MultipleProductResponse{
		Message:  message,
		Products: products,
	}
}
