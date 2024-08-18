package product

import (
	"context"
	"net/http"
	"time"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	h.logger.Infof("delete product request incoming: %v", productId)
	if productId == "" {
		h.logger.Error("missing productId")
		utils.WriteError(w, http.StatusBadRequest, "productId missing")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.productServiceClient.DeleteProduct(ctx, &productPb.ProductRequest{
		ProductId: &productId,
	})
	if err != nil {
		h.logger.Errorf("failed to delete product: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("delete product request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
