package gateway

import (
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/product"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/user"
	"go.uber.org/zap"
)

type Service struct {
	httpAddr       string
	logger         *zap.SugaredLogger
	userService    ports.UserService
	productService ports.ProductService
}

func NewGatewayService(
	httpAddr string,
	logger *zap.SugaredLogger,
	userService user.Service,
	productService product.Service,
) *Service {
	return &Service{
		httpAddr:       httpAddr,
		logger:         logger,
		userService:    userService,
		productService: productService,
	}
}
