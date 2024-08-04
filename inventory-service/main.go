package main

import (
	"log"
	"net"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/config"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/inventory-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/service"
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/store"
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

	conn, err := net.Listen("tcp", vars.INVENTORY_SERVICE_GRPC_ADDR)
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()

	mongoClient, err := database.ConnectToDB(vars.DB_CONNECTION_LINK)
	if err != nil {
		logger.Fatalf("Could not connect to database: %v", err)
	}
	defer database.DisconnectOfDB(mongoClient)

	inventoryStore := store.NewInventoryStore(mongoClient, vars.INVENTORY_SERVICE_DATABASE_NAME, vars.INVENTORY_SERVICE_DATABASE_COLLECTION)
	svc := service.NewInventoryService(vars.INVENTORY_SERVICE_GRPC_ADDR, *inventoryStore, logger)
	inventoryPb.RegisterInventoryServiceServer(grpServer, svc)

	logger.Infof("gRPC server started in port: %v", vars.INVENTORY_SERVICE_GRPC_ADDR)
	if err := grpServer.Serve(conn); err != nil {
		logger.Fatal(err.Error())
	}
}
