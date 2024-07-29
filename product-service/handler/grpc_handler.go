package handler

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/product-service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	productDto, err := createProductSchemaDto(in.GetProduct())
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	productId, err := database.GetMongoId(productDto.ID.Hex())
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	_, err = h.service.ProductStore.GetOne(ctx, productId)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return &productPb.ProductResponse{}, errors.New(utils.NOT_FOUND)
		}
	}

	createdProduct, err := h.service.ProductStore.CreateOne(ctx, productDto)
	if err != nil {
		return &productPb.ProductResponse{}, err
	}

	id, ok := createdProduct.InsertedID.(primitive.ObjectID)
	if !ok {
		return &productPb.ProductResponse{}, fmt.Errorf("failed to parse _id to string")
	}
	in.GetProduct().Id = id.Hex()

	return &productPb.ProductResponse{
		Message: "Product created successfully",
		Product: in.GetProduct(),
	}, nil
}
