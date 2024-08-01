package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	"github.com/Kiyosh31/ms-ecommerce/product-service/product_types"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProductService) CreateProduct(ctx context.Context, in *productPb.CreateProductRequest) (*productPb.ProductResponse, error) {
	s.logger.Infof("Create product incoming request: %v", in)

	productDto, err := createProductSchemaDto(in.GetProduct())
	if err != nil {
		s.logger.Errorf("error creating schema: %v", err)
		return &productPb.ProductResponse{}, err
	}

	productId, err := database.GetMongoId(productDto.ID.Hex())
	if err != nil {
		s.logger.Errorf("error creating productId: %v", err)
		return &productPb.ProductResponse{}, err
	}

	_, err = s.ProductStore.GetOne(ctx, productId)
	if err != nil {
		s.logger.Errorf("error getting product: %v", err)
		return &productPb.ProductResponse{}, err
	}

	createdProduct, err := s.ProductStore.CreateOne(ctx, productDto)
	if err != nil {
		s.logger.Errorf("error creating product: %v", err)
		return &productPb.ProductResponse{}, err
	}

	id, ok := createdProduct.InsertedID.(primitive.ObjectID)
	if !ok {
		s.logger.Errorf("error getting id: %v", err)
		return &productPb.ProductResponse{}, fmt.Errorf("failed to parse _id to string")
	}
	in.GetProduct().Id = id.Hex()

	s.logger.Infof("create product request finished: %v", createdProduct)
	return &productPb.ProductResponse{
		Message: "Product created successfully",
		Product: in.GetProduct(),
	}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, in *productPb.GetProductRequest) (*productPb.ProductResponse, error) {
	s.logger.Infof("Get product incoming request: %v", in)

	productId, err := database.GetMongoId(in.GetProductId())
	if err != nil {
		s.logger.Errorf("error getting productId: %v", err)
		return &productPb.ProductResponse{}, err
	}

	productFounded, err := s.ProductStore.GetOne(ctx, productId)
	if err != nil {
		s.logger.Errorf("error finding product: %v", err)
		return &productPb.ProductResponse{}, err
	}
	if reflect.DeepEqual(productFounded, product_types.ProductSchema{}) {
		s.logger.Errorf("error existing product: %v", err)
		return &productPb.ProductResponse{}, errors.New(utils.NOT_FOUND)
	}

	res := createProductResponseDto("Product founded", productFounded)

	s.logger.Infof("get product request finished: %v", res)
	return res, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context, in *productPb.GetAllProductsRequest) (*productPb.MultipleProductResponse, error) {
	s.logger.Infof("get all products incoming request: %v", in)

	productsFounded, err := s.ProductStore.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("error finding products: %v", err)
		return &productPb.MultipleProductResponse{}, err
	}

	res := createMultipleProductsResponseDto("products found", productsFounded)

	s.logger.Infof("get all products request finished: %v", res)
	return res, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, in *productPb.UpdateProductRequest) (*productPb.ProductResponse, error) {
	s.logger.Infof("Update product incoming request: %v", in)

	productId, err := database.GetMongoId(in.GetProduct().GetId())
	if err != nil {
		s.logger.Errorf("error getting productId: %v", err)
		return &productPb.ProductResponse{}, err
	}

	foundedProduct, err := s.ProductStore.GetOne(ctx, productId)
	if err != nil {
		s.logger.Errorf("error finding product: %v", err)
		return &productPb.ProductResponse{}, err
	}
	if reflect.DeepEqual(foundedProduct, product_types.ProductSchema{}) {
		s.logger.Errorf("error existing product: %v", err)
		return &productPb.ProductResponse{}, errors.New(utils.NOT_FOUND)
	}

	userToUpdate, err := createProductSchemaDto(in.GetProduct())
	if err != nil {
		s.logger.Errorf("error creating product schema: %v", err)
		return &productPb.ProductResponse{}, err
	}

	_, err = s.ProductStore.UpdateOne(ctx, userToUpdate)
	if err != nil {
		s.logger.Errorf("error updating product: %v", err)
		return &productPb.ProductResponse{}, err
	}

	res := createProductResponseDto("Product updated successfully", userToUpdate)

	s.logger.Infof("update product request finished: %v", err)
	return res, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, in *productPb.DeleteProductRequest) (*productPb.ProductResponse, error) {
	s.logger.Infof("Delete product incoming request: %v", in)

	productId, err := database.GetMongoId(in.GetProductId())
	if err != nil {
		s.logger.Errorf("error getting productId: %v", err)
		return &productPb.ProductResponse{}, err
	}

	foundedProduct, err := s.ProductStore.GetOne(ctx, productId)
	if err != nil {
		s.logger.Errorf("error finding product: %v", err)
		return &productPb.ProductResponse{}, err
	}
	if reflect.DeepEqual(foundedProduct, product_types.ProductSchema{}) {
		s.logger.Errorf("error existing product: %v", err)
		return &productPb.ProductResponse{}, errors.New(utils.NOT_FOUND)
	}

	_, err = s.ProductStore.DeleteOne(ctx, productId)
	if err != nil {
		s.logger.Errorf("error deleting product: %v", err)
		return &productPb.ProductResponse{}, err
	}

	res := createProductResponseDto("Product deleted successfully", foundedProduct)
	s.logger.Infof("product delete request finished: %v", res)
	return res, nil
}
