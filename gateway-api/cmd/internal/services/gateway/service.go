package gateway

import (
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/brand"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/category"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/product"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/user"
	"go.uber.org/zap"
)

type Service struct {
	httpAddr        string
	logger          *zap.SugaredLogger
	userService     ports.UserService
	productService  ports.ProductService
	brandService    ports.BrandService
	categoryService ports.CategoryService
}

func NewGatewayService(
	httpAddr string,
	logger *zap.SugaredLogger,
	userService *user.Service,
	productService *product.Service,
	brandService *brand.Service,
	categoryService *category.Service,
) *Service {
	return &Service{
		httpAddr:        httpAddr,
		logger:          logger,
		userService:     userService,
		productService:  productService,
		brandService:    brandService,
		categoryService: categoryService,
	}
}
