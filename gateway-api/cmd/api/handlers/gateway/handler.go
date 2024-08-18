package gateway

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/brand"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/category"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/product"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/user"
	"go.uber.org/zap"
)

type Handler struct {
	router          *http.ServeMux
	logger          *zap.SugaredLogger
	userHandler     user.Handler
	productHandler  product.Handler
	categoryHandler category.Handler
	brandHandler    brand.Handler
}

func NewGatewayHandler(
	router *http.ServeMux,
	logger *zap.SugaredLogger,
	userHandler user.Handler,
	productHandler product.Handler,
	categoryHandler category.Handler,
	brandHandler brand.Handler,
) *Handler {
	return &Handler{
		router:          router,
		logger:          logger,
		userHandler:     userHandler,
		productHandler:  productHandler,
		categoryHandler: categoryHandler,
		brandHandler:    brandHandler,
	}
}
