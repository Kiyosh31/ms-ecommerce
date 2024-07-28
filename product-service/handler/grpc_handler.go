package handler

import (
	"context"
	"log"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/product-service/service"
	"google.golang.org/grpc"
)

type ProductServiceGrpcHandler struct {
	productPb.UnimplementedProductServiceServer
	service service.ProductService
}

func NewGrpcProductServiceHandler(grpcServer *grpc.Server, service service.ProductService) {
	handler := &ProductServiceGrpcHandler{
		service: service,
	}

	productPb.RegisterProductServiceServer(grpcServer, handler)
}

func (h *ProductServiceGrpcHandler) CreateProduct(ctx context.Context, in *productPb.CreateProductRequest) (*productPb.ProductResponse, error) {
	log.Printf("Create product received request! %v", in)

	return &productPb.ProductResponse{
		Message: "yay",
	}, nil
}
