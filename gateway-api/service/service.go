package service

import (
	"log"
	"net/http"

	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayService struct {
	httpAddr           string
	userClientGrpcAddr string
}

func NewGatewayHttpService(httpAddr string, userClientGrpcAddr string) *GatewayService {
	return &GatewayService{
		httpAddr:           httpAddr,
		userClientGrpcAddr: userClientGrpcAddr,
	}
}

func (s *GatewayService) Run() {
	userServiceGrpcClient, conn := s.runUserServiceGrpcClient()
	defer conn.Close()

	mux := http.NewServeMux()
	handler := handler.NewHandler(userServiceGrpcClient)
	handler.RegisterRoutes(mux)

	log.Println("Http server starting at: ", s.httpAddr)

	if err := http.ListenAndServe(s.httpAddr, mux); err != nil {
		log.Fatalf("Failed to start http server: %v", err)
	}
}

func (s *GatewayService) runUserServiceGrpcClient() (userPb.UserServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.userClientGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Dialing user service at: ", s.userClientGrpcAddr)
	return userPb.NewUserServiceClient(conn), conn
}
