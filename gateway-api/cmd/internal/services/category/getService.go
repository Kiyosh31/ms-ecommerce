package category

import (
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s Service) GetService() (productPb.CategoryServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.categoryClientGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Fatalf("Failed to start server: %v", err)
	}

	s.logger.Infof("Dialing category-service at: %v", s.categoryClientGrpcAddr)
	return productPb.NewCategoryServiceClient(conn), conn
}
