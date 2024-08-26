package main

import (
	"log"
	"net"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/api/handler/payment"
	"github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/internal/config"
	paymentRepo "github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/internal/repositories/mongo/payment"
	paymentService "github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/internal/services/payment"
	paymentPb "github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/proto"
	"google.golang.org/grpc"
)

func main() {
	logger, err := customlogger.InitLogger()
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}
	defer logger.Sync()

	vars, err := config.LoadEnvVars()
	if err != nil {
		logger.Fatalf("Could not load env var: %v", err)
	}

	grpServer := grpc.NewServer()

	conn, err := net.Listen("tcp", vars.PAYMENT_SERVICE_GRPC_ADDR)
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()

	mongoClient, err := database.ConnectToDB(vars.DB_CONNECTION_LINK)
	if err != nil {
		logger.Fatalf("Could not connect to database: %v", err)
	}
	defer database.DisconnectOfDB(mongoClient)

	paymentRepo := paymentRepo.NewPaymentRepository(mongoClient, vars.PAYMENT_SERVICE_DATABASE_NAME, vars.PAYMENT_SERVICE_DATABASE_COLLECTION)
	paymentService := paymentService.NewUserService(paymentRepo, logger)
	paymentHandler := payment.NewPaymentHandler(paymentService, logger)
	paymentPb.RegisterPaymentServiceServer(grpServer, paymentHandler)

	logger.Infof("gRPC server started in port: %v", vars.PAYMENT_SERVICE_GRPC_ADDR)
	if err := grpServer.Serve(conn); err != nil {
		logger.Fatal(err.Error())
	}
}
