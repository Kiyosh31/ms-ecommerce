package payment

import (
	"github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/internal/services/payment"
	paymentPb "github.com/Kiyosh31/ms-ecommerce/payment-service/cmd/proto"
	"go.uber.org/zap"
)

type Handler struct {
	paymentPb.UnimplementedPaymentServiceServer
	paymentService ports.PaymentService
	logger         *zap.SugaredLogger
}

func NewPaymentHandler(paymentService *payment.Service, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		paymentService: paymentService,
		logger:         logger,
	}
}
