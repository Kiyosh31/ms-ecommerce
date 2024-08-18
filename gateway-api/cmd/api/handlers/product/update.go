package product

import (
	"context"
	"net/http"
	"time"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	h.logger.Infof("update product request incoming with id: %v and body: %v", productId, customlogger.ReadRequestPayload(r))

	var payload productPb.Product

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read product payload: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateProductPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate product payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.productServiceClient.UpdateProduct(ctx, &productPb.ProductRequest{
		ProductId: &productId,
		Product:   &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to update product: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("update product request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
