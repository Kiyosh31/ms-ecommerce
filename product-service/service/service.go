package service

import (
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/product-service/store"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
	GrpcAdd      string
	ProductStore store.ProductStore
	logger       *zap.SugaredLogger
	queue        amqp.Queue
	channel      *amqp.Channel
}

func NewProductService(
	grpcAdd string,
	productStore store.ProductStore,
	logger *zap.SugaredLogger,
	queue amqp.Queue,
	channel *amqp.Channel,
) *ProductService {
	return &ProductService{
		GrpcAdd:      grpcAdd,
		ProductStore: productStore,
		logger:       logger,
		queue:        queue,
		channel:      channel,
	}
}
