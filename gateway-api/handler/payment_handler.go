package handler

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	paymentPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/payment-service"
)

func (h *GatewayApiHandler) createPayment(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create payment request incoming: %v", customlogger.ReadRequestPayload(r))

	id := "123"
	res, err := h.paymentServiceGrpcClient.CreatePayment(r.Context(), &paymentPb.PaymentRequest{
		PaymentId: &id,
	})
	if err != nil {
		h.logger.Errorf("failed to create payment: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create payment finished: %v", res)
	utils.WriteResponse(w, http.StatusCreated, res)
}
