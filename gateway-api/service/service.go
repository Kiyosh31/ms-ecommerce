package service

import (
	"log"
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce/gateway-api/config"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/handler"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayService struct {
	addr string
}

func NewGatewayHttpService(addr string) *GatewayService {
	return &GatewayService{
		addr: addr,
	}
}

func (s *GatewayService) Run() {
	userServiceGrpcClient, conn := runUserServiceGrpcClient(config.GlobalEnvVars.USER_SERVICE_GRPC_ADDR)
	defer conn.Close()

	mux := http.NewServeMux()
	handler := handler.NewHandler(userServiceGrpcClient)
	handler.RegisterRoutes(mux)

	log.Println("Http server starting at: ", s.addr)

	if err := http.ListenAndServe(s.addr, mux); err != nil {
		log.Fatalf("Failed to start http server: %v", err)
	}
}

func runUserServiceGrpcClient(userServiceAddr string) (userPb.UserServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Dialing user service at: ", userServiceAddr)
	return userPb.NewUserServiceClient(conn), conn
}
