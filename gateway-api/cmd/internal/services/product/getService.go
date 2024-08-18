package product

import (
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s Service) GetService() (productPb.ProductServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.productClientGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Fatalf("Failed to start server: %v", err)
	}

	s.logger.Infof("Dialing user service at: %v", s.productClientGrpcAddr)
	return productPb.NewProductServiceClient(conn), conn
}
