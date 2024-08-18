package brand

import (
	"context"
	"net/http"
	"time"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
)

func (h *Handler) GetBrand(w http.ResponseWriter, r *http.Request) {
	brandId := r.PathValue("brandId")
	h.logger.Infof("get brand request incoming: %v", brandId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.brandServiceClient.GetBrand(ctx, &productPb.BrandRequest{
		BrandId: &brandId,
	})
	if err != nil {
		h.logger.Errorf("failed to get brand: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get brand finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)

}
