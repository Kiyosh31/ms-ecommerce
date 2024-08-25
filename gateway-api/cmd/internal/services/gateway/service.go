package gateway

import (
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/brand"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/category"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/inventory"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/product"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/user"
	"go.uber.org/zap"
)

type Service struct {
	httpAddr         string
	secretKey        string
	logger           *zap.SugaredLogger
	userService      ports.UserService
	productService   ports.ProductService
	brandService     ports.BrandService
	categoryService  ports.CategoryService
	inventoryService ports.InventoryService
}

func NewGatewayService(
	httpAddr string,
	secretKey string,
	logger *zap.SugaredLogger,
	userService *user.Service,
	productService *product.Service,
	brandService *brand.Service,
	categoryService *category.Service,
	inventoryService *inventory.Service,
) *Service {
	return &Service{
		httpAddr:         httpAddr,
		secretKey:        secretKey,
		logger:           logger,
		userService:      userService,
		productService:   productService,
		brandService:     brandService,
		categoryService:  categoryService,
		inventoryService: inventoryService,
	}
}
