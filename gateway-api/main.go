package main

import (
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/service"
)

func main() {
	svc := service.NewGatewayHttpService("localhost:3000")
	svc.Run()
}
