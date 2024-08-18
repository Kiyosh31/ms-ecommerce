package product

import (
	"context"
	"net/http"
	"time"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("get product request incoming: %v", customlogger.ReadRequestPayload(r))
	productId := r.PathValue("productId")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.productServiceClient.GetProduct(ctx, &productPb.ProductRequest{
		ProductId: &productId,
	})
	if err != nil {
		h.logger.Errorf("failed to get product: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get product request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
