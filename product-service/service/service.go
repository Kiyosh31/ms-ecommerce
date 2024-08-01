package service

import (
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/product-service/store"
	"go.uber.org/zap"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
	GrpcAdd      string
	ProductStore store.ProductStore
	logger       *zap.SugaredLogger
}

func NewProductService(grpcAdd string, productStore store.ProductStore, logger *zap.SugaredLogger) *ProductService {
	return &ProductService{
		GrpcAdd:      grpcAdd,
		ProductStore: productStore,
		logger:       logger,
	}
}
