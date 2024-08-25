package main

import (
	"log"
	"net"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/database"
	inventoryHandler "github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/api/handlers/inventory"
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/config"
	inventoryRepo "github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/repositories/mongo/inventory"
	inventoryService "github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/services/inventory"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/proto"
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

	inventoryRepository := inventoryRepo.NewInventoryRepository(
		mongoClient,
		vars.INVENTORY_SERVICE_DATABASE_NAME,
		vars.INVENTORY_SERVICE_DATABASE_COLLECTION,
	)
	inventoryService := inventoryService.NewInventoryService(
		inventoryRepository,
		logger,
	)
	inventoryHandler := inventoryHandler.NewInventoryHandler(
		inventoryService,
		logger,
	)
	inventoryPb.RegisterInventoryServiceServer(grpServer, inventoryHandler)

	logger.Infof("gRPC server started in port: %v", vars.INVENTORY_SERVICE_GRPC_ADDR)
	if err := grpServer.Serve(conn); err != nil {
		logger.Fatal(err.Error())
	}
}
