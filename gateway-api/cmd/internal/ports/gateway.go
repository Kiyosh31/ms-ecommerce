package ports

import (
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/inventory-service"
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

type BrandService interface {
	GetService() (productPb.BrandServiceClient, *grpc.ClientConn)
}

type CategoryService interface {
	GetService() (productPb.CategoryServiceClient, *grpc.ClientConn)
}

type InventoryService interface {
	GetService() (inventoryPb.InventoryServiceClient, *grpc.ClientConn)
}
