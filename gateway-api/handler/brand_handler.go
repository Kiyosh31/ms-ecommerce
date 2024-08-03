package handler

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
)

func (h *GatewayApiHandler) createBrand(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create brand request incoming: %v", customlogger.ReadRequestPayload(r))

	var payload productPb.Brand

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read brand payload: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateBrandPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate brand payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.productServiceClient.CreateBrand(r.Context(), &productPb.BrandRequest{
		Brand: &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to create category: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create brand finished: %v", res)
	utils.WriteResponse(w, http.StatusCreated, res)

}

func (h *GatewayApiHandler) getBrand(w http.ResponseWriter, r *http.Request) {
}

func (h *GatewayApiHandler) updateBrand(w http.ResponseWriter, r *http.Request) {
}

func (h *GatewayApiHandler) deleteBrand(w http.ResponseWriter, r *http.Request) {
}
