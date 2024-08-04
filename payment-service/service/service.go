package service

import (
	paymentPb "github.com/Kiyosh31/ms-ecommerce/payment-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/payment-service/store"
	"go.uber.org/zap"
)

type PaymentService struct {
	paymentPb.UnimplementedPaymentServiceServer
	GrpcAddr     string
	PaymentStore store.PaymentStore
	logger       *zap.SugaredLogger
}

func NewPaymentService(grpcAddr string, paymentStore store.PaymentStore, logger *zap.SugaredLogger) *PaymentService {
	return &PaymentService{
		GrpcAddr:     grpcAddr,
		PaymentStore: paymentStore,
		logger:       logger,
	}
}
