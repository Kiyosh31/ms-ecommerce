package service

import (
	"log"
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce/gateway-api/handler"
)

type GatewayServiceHttp struct {
	addr string
}

func NewGatewayHttpService(addr string) *GatewayServiceHttp {
	return &GatewayServiceHttp{
		addr: addr,
	}
}

func (s *GatewayServiceHttp) Run() {
	mux := http.NewServeMux()
	handler := handler.NewHandler()
	handler.RegisterRoutes(mux)

	log.Println("Http server starting at: ", s.addr)

	if err := http.ListenAndServe(s.addr, mux); err != nil {
		log.Fatalf("Failed to start http server: %v", err)
	}
}
