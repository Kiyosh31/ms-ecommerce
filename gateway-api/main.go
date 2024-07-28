package main

import (
	"log"

	"github.com/Kiyosh31/ms-ecommerce/gateway-api/config"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/service"
)

func main() {
	vars, err := config.LoadEnvVars()
	if err != nil {
		log.Fatalf("Could not load env vars: %v", err)
	}

	svc := service.NewGatewayHttpService(vars.GATEWAY_API_HTTP_ADRR, vars.GATEWAY_API_USER_SERVICE_GRPC_ADDR, vars.GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR)
	svc.Run()
}
