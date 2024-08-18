package brand

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
)

func (h *Handler) DeleteBrand(w http.ResponseWriter, r *http.Request) {
	brandId := r.PathValue("brandId")
	h.logger.Infof("delete brand request incoming with id: %v", brandId)

	res, err := h.brandServiceClient.DeleteBrand(r.Context(), &productPb.BrandRequest{
		BrandId: &brandId,
	})
	if err != nil {
		h.logger.Errorf("failed to delete brand: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("delete brand finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
