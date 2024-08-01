package handler

import (
	"net/http"

	cartPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/cart-service"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
	"go.uber.org/zap"
)

type GatewayApiHandler struct {
	userServiceClient    userPb.UserServiceClient
	productServiceClient productPb.ProductServiceClient
	cartProductClient    cartPb.CartServiceClient
	logger               *zap.SugaredLogger
}

func NewHandler(
	userServiceClient userPb.UserServiceClient,
	productServiceClient productPb.ProductServiceClient,
	cartProductClient cartPb.CartServiceClient,
	logger *zap.SugaredLogger,
) *GatewayApiHandler {
	return &GatewayApiHandler{
		userServiceClient:    userServiceClient,
		productServiceClient: productServiceClient,
		cartProductClient:    cartProductClient,
		logger:               logger,
	}
}

func (h *GatewayApiHandler) RegisterRoutes(mux *http.ServeMux) {
	// User endpoints
	mux.HandleFunc("POST /api/v1/user", h.createUser)
	mux.HandleFunc("GET /api/v1/user/{userId}", h.getUser)
	mux.HandleFunc("PUT /api/v1/user/{userId}", h.updateUser)
	mux.HandleFunc("DELETE /api/v1/user/{userId}", h.deleteUser)
	mux.HandleFunc("POST /api/v1/user/reactivate", h.reactivateUser)

	// Product endpoints
	mux.HandleFunc("POST /api/v1/product", h.createProduct)
	mux.HandleFunc("GET /api/v1/product/{productId}", h.getProduct)
	mux.HandleFunc("GET /api/v1/product/all", h.getAllProducts)
	mux.HandleFunc("PUT /api/v1/product/{productId}", h.updateProduct)
	mux.HandleFunc("DELETE /api/v1/product/{productId}", h.deleteProduct)

	// Cart endpoints
	mux.HandleFunc("POST /api/v1/cart", h.createCart)
	mux.HandleFunc("GET /api/v1/cart/{userId}/{cartId}", h.getCart)
	mux.HandleFunc("GET /api/v1/cart/all/{userId}", h.getAllCarts)

}
