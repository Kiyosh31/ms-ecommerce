package handler

import (
	"net/http"

	inventoryPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/inventory-service"
	orderPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/order-service"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
	"go.uber.org/zap"
)

type GatewayApiHandler struct {
	userServiceClient      userPb.UserServiceClient
	productServiceClient   productPb.ProductServiceClient
	inventoryServiceClient inventoryPb.InventoryServiceClient
	orderServiceGrpcClient orderPb.OrderServiceClient
	logger                 *zap.SugaredLogger
}

func NewHandler(
	userServiceClient userPb.UserServiceClient,
	productServiceClient productPb.ProductServiceClient,
	inventoryServiceClient inventoryPb.InventoryServiceClient,
	orderServiceGrpcClient orderPb.OrderServiceClient,
	logger *zap.SugaredLogger,
) *GatewayApiHandler {
	return &GatewayApiHandler{
		userServiceClient:      userServiceClient,
		productServiceClient:   productServiceClient,
		inventoryServiceClient: inventoryServiceClient,
		orderServiceGrpcClient: orderServiceGrpcClient,
		logger:                 logger,
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

	// Category endpoints
	router.HandleFunc("POST /api/v1/category", h.createCategory)
	router.HandleFunc("GET /api/v1/category/{categoryId}", h.getCategory)
	router.HandleFunc("PUT /api/v1/category/{categoryId}", h.updateCategory)
	router.HandleFunc("DELETE /api/v1/category/{categoryId}", h.deleteCategory)

	// Brand endpoints
	router.HandleFunc("POST /api/v1/brand", h.createBrand)
	router.HandleFunc("GET /api/v1/brand/{brandId}", h.getBrand)
	router.HandleFunc("PUT /api/v1/brand/{brandId}", h.updateBrand)
	router.HandleFunc("DELETE /api/v1/brand/{brandId}", h.deleteBrand)

	// inventory endpoints
	router.HandleFunc("POST /api/v1/inventory", h.createInventory)
	router.HandleFunc("GET /api/v1/inventory/{inventoryId}", h.getInventory)
	router.HandleFunc("PUT /api/v1/inventory/{inventoryId}", h.updateInventory)

	// order endpoints
	router.HandleFunc("POST /api/v1/order", h.createOrder)
	router.HandleFunc("GET /api/v1/order/{orderId}", h.getOrder)
	router.HandleFunc("POST /api/v1/order/{orderId}/cancel", h.cancelOrder)
}
