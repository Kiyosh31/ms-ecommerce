package product

import (
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/services/product"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
	"go.uber.org/zap"
)

type Handler struct {
	productPb.UnimplementedProductServiceServer
	productService ports.ProductService
	logger         *zap.SugaredLogger
}

func NewProductHandler(
	productService *product.Service,
	logger *zap.SugaredLogger,
) *Handler {
	return &Handler{
		productService: productService,
		logger:         logger,
	}
}
