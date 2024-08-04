package handler

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
)

func (h *GatewayApiHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create order request incoming: %v", customlogger.ReadRequestPayload(r))
}

func (h *GatewayApiHandler) getOrder(w http.ResponseWriter, r *http.Request) {

}

func (h *GatewayApiHandler) cancelOrder(w http.ResponseWriter, r *http.Request) {

}
