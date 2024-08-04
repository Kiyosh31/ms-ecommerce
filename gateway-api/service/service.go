package service

import (
	"net/http"

	inventoryPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/inventory-service"
	orderPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/order-service"
	paymentPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/payment-service"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/handler"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayService struct {
	httpAddr                string
	userClientGrpcAddr      string
	productClientGrpcAddr   string
	inventoryClientGrpcAddr string
	orderClientGrpcAdd      string
	paymentClientGrpcAdd    string
	logger                  *zap.SugaredLogger
}

func NewGatewayHttpService(
	httpAddr string,
	userClientGrpcAddr string,
	productClientGrpcAddr string,
	inventoryClientGrpcAddr string,
	orderClientGrpcAdd string,
	paymentClientGrpcAdd string,
	logger *zap.SugaredLogger,
) *GatewayService {
	return &GatewayService{
		httpAddr:                httpAddr,
		userClientGrpcAddr:      userClientGrpcAddr,
		productClientGrpcAddr:   productClientGrpcAddr,
		inventoryClientGrpcAddr: inventoryClientGrpcAddr,
		orderClientGrpcAdd:      orderClientGrpcAdd,
		paymentClientGrpcAdd:    paymentClientGrpcAdd,
		logger:                  logger,
	}
}

func (s *GatewayService) Run() {
	userServiceGrpcClient, conn := s.runUserServiceGrpcClient()
	defer conn.Close()

	productServiceGrpcClient, productConn := s.runProductServiceGrpcClient()
	defer productConn.Close()

	inventoryServiceGrpcClient, inventoryConn := s.runInventoryServiceGrpcClient()
	defer inventoryConn.Close()

	orderServiceGrpcClient, orderConn := s.runOrderServiceGrpcClient()
	defer orderConn.Close()

	paymentServiceGrpcClient, paymentConn := s.runPaymentServiceGrpcClient()
	defer paymentConn.Close()

	handler := handler.NewHandler(
		userServiceGrpcClient,
		productServiceGrpcClient,
		inventoryServiceGrpcClient,
		orderServiceGrpcClient,
		paymentServiceGrpcClient,
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

func (s *GatewayService) runInventoryServiceGrpcClient() (inventoryPb.InventoryServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.inventoryClientGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Fatalf("Failed to start server: %v", err)
	}

	s.logger.Infof("Dialing product service at: %v", s.inventoryClientGrpcAddr)
	return inventoryPb.NewInventoryServiceClient(conn), conn
}

func (s *GatewayService) runOrderServiceGrpcClient() (orderPb.OrderServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.orderClientGrpcAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Fatalf("Failed to start server: %v", err)
	}

	s.logger.Infof("Dialing product service at: %v", s.orderClientGrpcAdd)
	return orderPb.NewOrderServiceClient(conn), conn
}

func (s *GatewayService) runPaymentServiceGrpcClient() (paymentPb.PaymentServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.paymentClientGrpcAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Fatalf("Failed to start server: %v", err)
	}

	s.logger.Infof("Dialing product service at: %v", s.paymentClientGrpcAdd)
	return paymentPb.NewPaymentServiceClient(conn), conn
}
