package inventory

import (
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/inventory-service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s Service) GetService() (inventoryPb.InventoryServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.inventoryClientGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Fatalf("Failed to start server: %v", err)
	}

	s.logger.Infof("Dialing inventory-service at: %v", s.inventoryClientGrpcAddr)
	return inventoryPb.NewInventoryServiceClient(conn), conn
}
