package service

import (
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/product-service/store"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
	GrpcAdd      string
	ProductStore store.ProductStore
}

func NewProductService(grpcAdd string, productStore store.ProductStore) *ProductService {
	return &ProductService{
		GrpcAdd:      grpcAdd,
		ProductStore: productStore,
	}
}
