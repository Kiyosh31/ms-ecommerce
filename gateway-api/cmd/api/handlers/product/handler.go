package product

import (
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
	"go.uber.org/zap"
)

type Handler struct {
	productServiceClient productPb.ProductServiceClient
	logger               *zap.SugaredLogger
}

func NewProductHandler(productServiceClient productPb.ProductServiceClient,
	logger *zap.SugaredLogger) *Handler {
	return &Handler{
		productServiceClient: productServiceClient,
		logger:               logger,
	}
}
