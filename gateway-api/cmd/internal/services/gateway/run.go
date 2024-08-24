package gateway

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/brand"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/category"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/gateway"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/product"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/user"
)

func (s Service) Run() {
	userServiceClient, userServiceConn := s.userService.GetService()
	defer userServiceConn.Close()
	userHandler := user.NewUserHandler(userServiceClient, s.logger)

	productServiceClient, productServiceConn := s.productService.GetService()
	defer productServiceConn.Close()
	productHandler := product.NewProductHandler(productServiceClient, s.logger)

	categoryServiceClient, categoryServiceConn := s.categoryService.GetService()
	defer categoryServiceConn.Close()
	categoryHandler := category.NewCategoryHandler(categoryServiceClient, s.logger)

	brandServiceClient, brandServiceConn := s.brandService.GetService()
	defer brandServiceConn.Close()
	brandHandler := brand.NewBrandHandler(brandServiceClient, s.logger)

	router := http.NewServeMux()
	gatewayHandler := gateway.NewGatewayHandler(
		router,
		s.logger,
		*userHandler,
		*productHandler,
		*categoryHandler,
		*brandHandler,
	)
	gatewayHandler.RegisterRoutes()

	s.logger.Infof("Http server starting at: %v", s.httpAddr)
	if err := http.ListenAndServe(s.httpAddr, router); err != nil {
		s.logger.Fatalf("Failed to start http server: %v", err)
	}
}
