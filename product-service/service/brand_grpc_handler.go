package service

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
)

func (s *ProductService) CreateBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	return &productPb.BrandResponse{}, nil
}

func (s *ProductService) GetBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	return &productPb.BrandResponse{}, nil

}

func (s *ProductService) UpdateBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	return &productPb.BrandResponse{}, nil

}

func (s *ProductService) DeleteBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	return &productPb.BrandResponse{}, nil

}
