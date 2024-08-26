package main

import (
	"log"
	"net"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/order-service/cmd/api/handler/order"
	"github.com/Kiyosh31/ms-ecommerce/order-service/cmd/internal/config"
	orderRepo "github.com/Kiyosh31/ms-ecommerce/order-service/cmd/internal/repositories/mongo/order"
	orderService "github.com/Kiyosh31/ms-ecommerce/order-service/cmd/internal/services/order"
	orderPb "github.com/Kiyosh31/ms-ecommerce/order-service/cmd/proto"
	"google.golang.org/grpc"
)

func main() {
	logger, err := customlogger.InitLogger()
	if err != nil {
		log.Fatalf("error logger init: %v", err)
	}

	vars, err := config.LoadEnvVars()
	if err != nil {
		logger.Fatalf("Could not load env var: %v", err)
	}

	grpServer := grpc.NewServer()

	conn, err := net.Listen("tcp", vars.ORDER_SERVICE_GRPC_ADDR)
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()

	mongoClient, err := database.ConnectToDB(vars.DB_CONNECTION_LINK)
	if err != nil {
		logger.Fatalf("Could not connect to database: %v", err)
	}
	defer database.DisconnectOfDB(mongoClient)

	orderRepository := orderRepo.NewOrderRepository(mongoClient, vars.ORDER_SERVICE_DATABASE_NAME, vars.ORDER_SERVICE_DATABASE_COLLECTION)
	orderService := orderService.NewOrderService(orderRepository, logger)
	orderHandler := order.NewOrderHandler(orderService, logger)
	orderPb.RegisterOrderServiceServer(grpServer, orderHandler)

	logger.Infof("gRPC server started in port: %v", vars.ORDER_SERVICE_GRPC_ADDR)
	if err := grpServer.Serve(conn); err != nil {
		logger.Fatal(err.Error())
	}
}
