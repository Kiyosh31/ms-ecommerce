package brand

import (
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/services/brand"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
	"go.uber.org/zap"
)

type Handler struct {
	productPb.UnimplementedBrandServiceServer
	brandService ports.BrandService
	logger       *zap.SugaredLogger
}

func NewBrandHandler(
	brandService *brand.Service,
	logger *zap.SugaredLogger,
) *Handler {
	return &Handler{
		brandService: brandService,
		logger:       logger,
	}
}
