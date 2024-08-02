package service

import (
	"net/http"

	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/handler"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayService struct {
	httpAddr              string
	userClientGrpcAddr    string
	productClientGrpcAddr string
	logger                *zap.SugaredLogger
}

func NewGatewayHttpService(
	httpAddr string,
	userClientGrpcAddr string,
	productClientGrpcAddr string,
	logger *zap.SugaredLogger,
) *GatewayService {
	return &GatewayService{
		httpAddr:              httpAddr,
		userClientGrpcAddr:    userClientGrpcAddr,
		productClientGrpcAddr: productClientGrpcAddr,
		logger:                logger,
	}
}

func (s *GatewayService) Run() {
	userServiceGrpcClient, conn := s.runUserServiceGrpcClient()
	defer conn.Close()

	productServiceGrpcClient, productConn := s.runProductServiceGrpcClient()
	defer productConn.Close()

	handler := handler.NewHandler(
		userServiceGrpcClient,
		productServiceGrpcClient,
		s.logger,
	)

	router := http.NewServeMux()
	handler.RegisterRoutes(router)

	s.logger.Infof("Http server starting at: %v", s.httpAddr)

	if err := http.ListenAndServe(s.httpAddr, router); err != nil {
		s.logger.Fatalf("Failed to start http server: %v", err)
	}
}

func (s *GatewayService) runUserServiceGrpcClient() (userPb.UserServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.userClientGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Fatalf("Failed to start server: %v", err)
	}

	s.logger.Infof("Dialing user service at: %v", s.userClientGrpcAddr)
	return userPb.NewUserServiceClient(conn), conn
}

func (s *GatewayService) runProductServiceGrpcClient() (productPb.ProductServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.productClientGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Fatalf("Failed to start server: %v", err)
	}

	s.logger.Infof("Dialing product service at: %v", s.productClientGrpcAddr)
	return productPb.NewProductServiceClient(conn), conn
}
