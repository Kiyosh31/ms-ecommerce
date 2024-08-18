package ports

import (
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
	"google.golang.org/grpc"
)

type GatewayService interface {
	Run()
}

type UserService interface {
	GetService() (userPb.UserServiceClient, *grpc.ClientConn)
}

type ProductService interface {
	GetService() (productPb.ProductServiceClient, *grpc.ClientConn)
}
