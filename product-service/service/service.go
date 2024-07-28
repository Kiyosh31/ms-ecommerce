package service

import "github.com/Kiyosh31/ms-ecommerce/product-service/store"

type ProductService struct {
	GrpcAdd      string
	ProductStore store.ProductStore
}

func NewProductService(grpcAdd string, productStore store.ProductStore) *ProductService {
	return &ProductService{
		GrpcAdd:      grpcAdd,
		ProductStore: productStore,
	}
}
