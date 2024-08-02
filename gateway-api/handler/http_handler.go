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

func (h *GatewayApiHandler) RegisterRoutes(router *http.ServeMux) {
	// User endpoints
	router.HandleFunc("POST /api/v1/user", h.createUser)
	router.HandleFunc("GET /api/v1/user/{userId}", h.getUser)
	router.HandleFunc("PUT /api/v1/user/{userId}", h.updateUser)
	router.HandleFunc("DELETE /api/v1/user/{userId}", h.deleteUser)
	router.HandleFunc("POST /api/v1/user/reactivate", h.reactivateUser)

	// Product endpoints
	router.HandleFunc("POST /api/v1/product", h.createProduct)
	router.HandleFunc("GET /api/v1/product/{productId}", h.getProduct)
	router.HandleFunc("GET /api/v1/product/all", h.getAllProducts)
	router.HandleFunc("PUT /api/v1/product/{productId}", h.updateProduct)
	router.HandleFunc("DELETE /api/v1/product/{productId}", h.deleteProduct)

	// Cart endpoints
	router.HandleFunc("POST /api/v1/cart", h.createCart)
	router.HandleFunc("GET /api/v1/cart/{userId}/{cartId}", h.getCart)
	router.HandleFunc("GET /api/v1/cart/all/{userId}", h.getAllCarts)
}
