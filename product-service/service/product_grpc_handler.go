package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProductService) CreateProduct(ctx context.Context, in *productPb.ProductRequest) (*productPb.ProductResponse, error) {
	s.logger.Infof("Create product incoming request: %v", in)

	categoryId, err := database.GetMongoId(in.GetProduct().GetCategory().GetId())
	if err != nil {
		s.logger.Errorf("error getting categoryId: %v", err)
		return &productPb.ProductResponse{}, err
	}

	brandId, err := database.GetMongoId(in.GetProduct().GetBrand().GetId())
	if err != nil {
		s.logger.Errorf("error getting brandId: %v", err)
		return &productPb.ProductResponse{}, err
	}

	// check category exists
	_, exists, err := s.CategoryStore.CategoryExists(ctx, categoryId)
	if err != nil {
		s.logger.Errorf("error finding existing category: %v", err)
		return &productPb.ProductResponse{}, err
	}
	if !exists {
		s.logger.Errorf("error category dont exists: %v", err)
		return &productPb.ProductResponse{}, errors.New("category dont exist")
	}

	// check brand exists
	_, exists, err = s.BrandStore.BrandExists(ctx, brandId)
	if err != nil {
		s.logger.Errorf("error finding existing brand: %v", err)
		return &productPb.ProductResponse{}, err
	}
	if !exists {
		s.logger.Errorf("error category dont exists: %v", err)
		return &productPb.ProductResponse{}, errors.New("brand dont exist")
	}

	// create product schema for db
	productDto, err := createProductSchemaDto(in.GetProduct())
	if err != nil {
		s.logger.Errorf("error creating schema: %v", err)
		return &productPb.ProductResponse{}, err
	}

	// saving product to db
	createdProduct, err := s.ProductStore.CreateOne(ctx, productDto)
	if err != nil {
		s.logger.Errorf("error creating product: %v", err)
		return &productPb.ProductResponse{}, err
	}

	// getting productID created in db
	id, ok := createdProduct.InsertedID.(primitive.ObjectID)
	if !ok {
		s.logger.Errorf("error getting id: %v", err)
		return &productPb.ProductResponse{}, fmt.Errorf("failed to parse _id to string")
	}
	in.GetProduct().Id = id.Hex()

	// return response
	s.logger.Infof("create product request finished: %v", createdProduct)
	return &productPb.ProductResponse{
		Message: "Product created successfully",
		Product: []*productPb.Product{in.GetProduct()},
	}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, in *productPb.ProductRequest) (*productPb.ProductResponse, error) {
	s.logger.Infof("Get product incoming request: %v", in)

	productId, err := database.GetMongoId(in.GetProductId())
	if err != nil {
		s.logger.Errorf("error getting productId: %v", err)
		return &productPb.ProductResponse{}, err
	}

	productFounded, exists, err := s.ProductStore.ProductExists(ctx, productId)
	if err != nil {
		s.logger.Errorf("error finding existing product: %v", err)
		return &productPb.ProductResponse{}, err
	}
	if !exists {
		s.logger.Errorf("error product not exists: %v", err)
		return &productPb.ProductResponse{}, errors.New("product not exists")
	}

	res := createProductResponseDto("Product founded", productFounded)

	s.logger.Infof("get product request finished: %v", res)
	return res, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context, in *productPb.ProductRequest) (*productPb.ProductResponse, error) {
	s.logger.Infof("get all products incoming request: %v", in)

	productsFounded, err := s.ProductStore.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("error finding products: %v", err)
		return &productPb.ProductResponse{}, err
	}

	res := createMultipleProductsResponseDto("products found", productsFounded)

	s.logger.Infof("get all products request finished: %v", res)
	return res, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, in *productPb.ProductRequest) (*productPb.ProductResponse, error) {
	s.logger.Infof("Update product incoming request: %v", in)

	productId, err := database.GetMongoId(in.GetProductId())
	if err != nil {
		s.logger.Errorf("error getting productId: %v", err)
		return &productPb.ProductResponse{}, err
	}

	_, exists, err := s.ProductStore.ProductExists(ctx, productId)
	if err != nil {
		s.logger.Errorf("error finding existing product: %v", err)
		return &productPb.ProductResponse{}, err
	}
	if !exists {
		s.logger.Errorf("error product dont exists: %v", err)
		return &productPb.ProductResponse{}, errors.New("product not exist")
	}

	productToUpdate, err := createProductSchemaDto(in.GetProduct())
	if err != nil {
		s.logger.Errorf("error creating product schema: %v", err)
		return &productPb.ProductResponse{}, err
	}

	_, err = s.ProductStore.UpdateOne(ctx, productToUpdate)
	if err != nil {
		s.logger.Errorf("error updating product: %v", err)
		return &productPb.ProductResponse{}, err
	}

	res := createProductResponseDto("Product updated successfully", productToUpdate)

	s.logger.Infof("update product request finished: %v", err)
	return res, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, in *productPb.ProductRequest) (*productPb.ProductResponse, error) {
	s.logger.Infof("Delete product incoming request: %v", in)

	productId, err := database.GetMongoId(in.GetProductId())
	if err != nil {
		s.logger.Errorf("error getting productId: %v", err)
		return &productPb.ProductResponse{}, err
	}

	foundedProduct, exists, err := s.ProductStore.ProductExists(ctx, productId)
	if err != nil {
		s.logger.Errorf("error finding existing product: %v", err)
		return &productPb.ProductResponse{}, err
	}
	if !exists {
		s.logger.Errorf("error product dont exists: %v", err)
		return &productPb.ProductResponse{}, errors.New("product not exist")
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
