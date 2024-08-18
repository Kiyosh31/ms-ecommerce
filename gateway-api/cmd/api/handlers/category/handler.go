package category

import (
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
	"go.uber.org/zap"
)

type Handler struct {
	categoryServiceClient productPb.ProductServiceClient
	logger                *zap.SugaredLogger
}

func NewCategoryHandler(categoryServiceClient productPb.ProductServiceClient,
	logger *zap.SugaredLogger) *Handler {
	return &Handler{
		categoryServiceClient: categoryServiceClient,
		logger:                logger,
	}
}
