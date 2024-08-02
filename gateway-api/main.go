package main

import (
	"log"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/config"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/service"
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

	svc := service.NewGatewayHttpService(
		vars.GATEWAY_API_HTTP_ADRR,
		vars.GATEWAY_API_USER_SERVICE_GRPC_ADDR,
		vars.GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR,
		logger,
	)
	svc.Run()
}
