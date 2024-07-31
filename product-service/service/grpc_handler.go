package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	"github.com/Kiyosh31/ms-ecommerce/product-service/product_types"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return &productPb.ProductResponse{}, err
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

func (s *ProductService) GetProduct(ctx context.Context, in *productPb.GetProductRequest) (*productPb.ProductResponse, error) {
	log.Printf("Get product received request! %v", in)

	productId, err := database.GetMongoId(in.GetProductId())
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	productFounded, err := s.ProductStore.GetOne(ctx, productId)
	if err != nil {
		return &productPb.ProductResponse{}, err
	}
	if reflect.DeepEqual(productFounded, product_types.ProductSchema{}) {
		return &productPb.ProductResponse{}, errors.New(utils.NOT_FOUND)
	}

	res := createProductResponseDto("Product founded", productFounded)

	return res, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context, in *productPb.GetAllProductsRequest) (*productPb.MultipleProductResponse, error) {
	log.Printf("Get all products received request! %v", in)

	productsFounded, err := s.ProductStore.GetAll(ctx)
	if err != nil {
		return &productPb.MultipleProductResponse{}, err
	}

	res := createMultipleProductsResponseDto("products found", productsFounded)

	return res, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, in *productPb.UpdateProductRequest) (*productPb.ProductResponse, error) {
	log.Printf("Update product received request! %v", in)

	productId, err := database.GetMongoId(in.GetProduct().GetId())
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	foundedProduct, err := s.ProductStore.GetOne(ctx, productId)
	if err != nil {
		return &productPb.ProductResponse{}, err
	}
	if reflect.DeepEqual(foundedProduct, product_types.ProductSchema{}) {
		return &productPb.ProductResponse{}, errors.New(utils.NOT_FOUND)
	}

	userToUpdate, err := createProductSchemaDto(in.GetProduct())
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	_, err = s.ProductStore.UpdateOne(ctx, userToUpdate)
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	res := createProductResponseDto("Product updated successfully", userToUpdate)

	return res, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, in *productPb.DeleteProductRequest) (*productPb.ProductResponse, error) {
	log.Printf("Delete product received request! %v", in)

	productId, err := database.GetMongoId(in.GetProductId())
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	foundedProduct, err := s.ProductStore.GetOne(ctx, productId)
	if err != nil {
		return &productPb.ProductResponse{}, err
	}
	if reflect.DeepEqual(foundedProduct, product_types.ProductSchema{}) {
		return &productPb.ProductResponse{}, errors.New(utils.NOT_FOUND)
	}

	_, err = s.ProductStore.DeleteOne(ctx, productId)
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	res := createProductResponseDto("Product deleted successfully", foundedProduct)
	return res, nil
}
