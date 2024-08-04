package service

import (
	"context"

	paymentPb "github.com/Kiyosh31/ms-ecommerce/payment-service/proto"
)

func (s *PaymentService) CreatePayment(ctx context.Context, in *paymentPb.PaymentRequest) (*paymentPb.PaymentResponse, error) {
	s.logger.Infof("create payment request incoming: %v", in)
	return &paymentPb.PaymentResponse{}, nil
}
