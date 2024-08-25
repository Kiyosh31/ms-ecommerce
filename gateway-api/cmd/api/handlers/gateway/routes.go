package gateway

func (h *Handler) RegisterRoutes() {
	// User endpoints
	h.router.HandleFunc("POST /api/v1/user", h.userHandler.Create)
	h.router.HandleFunc("GET /api/v1/user/{userId}", h.userHandler.Get)
	h.router.HandleFunc("PUT /api/v1/user/{userId}", h.userHandler.Update)
	h.router.HandleFunc("DELETE /api/v1/user/{userId}", h.userHandler.Deactivate)
	h.router.HandleFunc("POST /api/v1/user/reactivate", h.userHandler.Reactivate)
	h.router.HandleFunc("POST /api/v1/user/login", h.userHandler.Login)

	// Product endpoints
	h.router.HandleFunc("POST /api/v1/product", h.productHandler.CreateProduct)
	h.router.HandleFunc("GET /api/v1/product/{productId}", h.productHandler.GetProduct)
	h.router.HandleFunc("PUT /api/v1/product/{productId}", h.productHandler.UpdateProduct)
	h.router.HandleFunc("DELETE /api/v1/product/{productId}", h.productHandler.DeleteProduct)

	// Category endpoints
	h.router.HandleFunc("POST /api/v1/category", h.categoryHandler.CreateCategory)
	h.router.HandleFunc("GET /api/v1/category/{categoryId}", h.categoryHandler.GetCategory)
	h.router.HandleFunc("PUT /api/v1/category/{categoryId}", h.categoryHandler.UpdateCategory)
	h.router.HandleFunc("DELETE /api/v1/category/{categoryId}", h.categoryHandler.DeleteCategory)

	// Brand endpoints
	h.router.HandleFunc("POST /api/v1/brand", h.brandHandler.CreateBrand)
	h.router.HandleFunc("GET /api/v1/brand/{brandId}", h.brandHandler.GetBrand)
	h.router.HandleFunc("PUT /api/v1/brand/{brandId}", h.brandHandler.UpdateBrand)
	h.router.HandleFunc("DELETE /api/v1/brand/{brandId}", h.brandHandler.DeleteBrand)

	// inventory endpoints
	h.router.HandleFunc("GET /api/v1/inventory/{inventoryId}", h.inventoryHandler.Get)

	// order endpoints
	// router.HandleFunc("POST /api/v1/order", h.createOrder)
	// router.HandleFunc("GET /api/v1/order/{orderId}", h.getOrder)
	// router.HandleFunc("POST /api/v1/order/{orderId}/cancel", h.cancelOrder)

	// payment endpoints
	// router.HandleFunc("POST /api/v1/payment", h.createPayment)
}
