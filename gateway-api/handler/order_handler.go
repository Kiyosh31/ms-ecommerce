package handler

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	orderPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/order-service"
)

func (h *GatewayApiHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create order request incoming: %v", customlogger.ReadRequestPayload(r))

	id := "123"
	res, err := h.orderServiceGrpcClient.CreateOrder(r.Context(), &orderPb.OrderRequest{
		OrderId: &id,
	})
	if err != nil {
		h.logger.Errorf("failed to create product: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create order finished: %v", res)
	utils.WriteResponse(w, http.StatusCreated, res)
}

func (h *GatewayApiHandler) getOrder(w http.ResponseWriter, r *http.Request) {

}

func (h *GatewayApiHandler) cancelOrder(w http.ResponseWriter, r *http.Request) {

}
