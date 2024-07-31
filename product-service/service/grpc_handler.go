package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ProductService) CreateProduct(ctx context.Context, in *productPb.CreateProductRequest) (*productPb.ProductResponse, error) {
	log.Printf("Create product received request! %v", in)

	productDto, err := createProductSchemaDto(in.GetProduct())
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	productId, err := database.GetMongoId(productDto.ID.Hex())
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	_, err = s.ProductStore.GetOne(ctx, productId)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return &productPb.ProductResponse{}, errors.New(utils.NOT_FOUND)
		}
	}

	createdProduct, err := s.ProductStore.CreateOne(ctx, productDto)
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	id, ok := createdProduct.InsertedID.(primitive.ObjectID)
	if !ok {
		return &productPb.ProductResponse{}, fmt.Errorf("failed to parse _id to string")
	}
	in.GetProduct().Id = id.Hex()

	return &productPb.ProductResponse{
		Message: "Product created successfully",
		Product: in.GetProduct(),
	}, nil
}
