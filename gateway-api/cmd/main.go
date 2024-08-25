package main

import (
	"log"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"

	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/config"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/brand"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/category"
	gatewayService "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/gateway"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/inventory"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/product"
	userService "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/internal/services/user"
)

func main() {
	logger, err := customlogger.InitLogger()
	if err != nil {
		log.Fatalf("could not init logger: %v", err)
	}

	vars, err := config.LoadEnvVars()
	if err != nil {
		logger.Fatalf("Could not load env vars: %v", err)
	}

	userService := userService.NewUserService(vars.GATEWAY_API_USER_SERVICE_GRPC_ADDR, logger)
	productService := product.NewProductService(vars.GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR, logger)
	categoryService := category.NewCategoryService(vars.GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR, logger)
	brandService := brand.NewBrandService(vars.GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR, logger)
	inventoryHandler := inventory.NewInventoryService(vars.GATEWAY_API_INVENTORY_SERVICE_GRPC_ADDR, logger)

	gatewayService := gatewayService.NewGatewayService(
		vars.GATEWAY_API_HTTP_ADRR,
		vars.SECRET_KEY,
		logger,
		userService,
		productService,
		brandService,
		categoryService,
		inventoryHandler,
	)
	gatewayService.Run()
}
