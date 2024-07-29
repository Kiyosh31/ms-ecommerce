package handler

import (
	"net/http"

	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
)

type GatewayApiHandler struct {
	userServiceClient    userPb.UserServiceClient
	productServiceClient productPb.ProductServiceClient
}

func NewHandler(userServiceClient userPb.UserServiceClient, productServiceClient productPb.ProductServiceClient) *GatewayApiHandler {
	return &GatewayApiHandler{
		userServiceClient:    userServiceClient,
		productServiceClient: productServiceClient,
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
}
