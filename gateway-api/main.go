package main

import (
	"fmt"
	"os"

	"github.com/Kiyosh31/ms-ecommerce/gateway-api/config"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/service"
)

func main() {
	errors := config.LoadEnvVars()
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Println("Error loading environment variables:", err)
		}
		os.Exit(1)
	}

	svc := service.NewGatewayHttpService(config.GlobalEnvVars.GATEWAY_API_HTTP_ADRR)
	svc.Run()
}
