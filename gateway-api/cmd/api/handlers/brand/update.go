package brand

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/product"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
)

func (h *Handler) UpdateBrand(w http.ResponseWriter, r *http.Request) {
	brandId := r.PathValue("brandId")
	h.logger.Infof("update brand request incoming with id: %v and payload: %v", brandId, customlogger.ReadRequestPayload(r))

	var payload productPb.Brand

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read brand payload: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := product.ValidateBrandPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate brand payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.brandServiceClient.UpdateBrand(r.Context(), &productPb.BrandRequest{
		BrandId: &brandId,
		Brand:   &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to update category: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("update brand finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
