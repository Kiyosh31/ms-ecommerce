package service

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
)

func (s *ProductService) CreateCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	return &productPb.CategoryResponse{}, nil
}

func (s *ProductService) GetCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	return &productPb.CategoryResponse{}, nil

}

func (s *ProductService) UpdateCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	return &productPb.CategoryResponse{}, nil

}

func (s *ProductService) DeleteCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	return &productPb.CategoryResponse{}, nil

}
