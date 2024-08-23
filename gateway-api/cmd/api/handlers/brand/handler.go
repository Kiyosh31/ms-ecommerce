package brand

import (
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
	"go.uber.org/zap"
)

type Handler struct {
	brandServiceClient productPb.BrandServiceClient
	logger             *zap.SugaredLogger
}

func NewBrandHandler(brandServiceClient productPb.BrandServiceClient,
	logger *zap.SugaredLogger) *Handler {
	return &Handler{
		brandServiceClient: brandServiceClient,
		logger:             logger,
	}
}
