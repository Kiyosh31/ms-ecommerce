package service

import (
	"log"
	"net/http"

	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayService struct {
	httpAddr              string
	userClientGrpcAddr    string
	productClientGrpcAddr string
}

func NewGatewayHttpService(
	httpAddr string,
	userClientGrpcAddr string,
	productClientGrpcAddr string,
) *GatewayService {
	return &GatewayService{
		httpAddr:              httpAddr,
		userClientGrpcAddr:    userClientGrpcAddr,
		productClientGrpcAddr: productClientGrpcAddr,
	}
}

func (s *GatewayService) Run() {
	userServiceGrpcClient, conn := s.runUserServiceGrpcClient()
	defer conn.Close()
	productServiceGrpcClient, productConn := s.runProductServiceGrpcClient()
	defer productConn.Close()

	mux := http.NewServeMux()
	handler := handler.NewHandler(userServiceGrpcClient, productServiceGrpcClient)
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

func (s *GatewayService) runProductServiceGrpcClient() (productPb.ProductServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.productClientGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Dialing product service at: ", s.productClientGrpcAddr)
	return productPb.NewProductServiceClient(conn), conn
}
