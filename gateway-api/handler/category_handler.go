package handler

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
)

func (h *GatewayApiHandler) createCategory(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create category request incoming: %v", customlogger.ReadRequestPayload(r))

	var payload productPb.Category

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read category payload: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateCategoryPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate category payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.productServiceClient.CreateCategory(r.Context(), &productPb.CategoryRequest{
		Category: &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to create category: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create category finished: %v", res)
	utils.WriteResponse(w, http.StatusCreated, res)
}

func (h *GatewayApiHandler) getCategory(w http.ResponseWriter, r *http.Request) {
}

func (h *GatewayApiHandler) updateCategory(w http.ResponseWriter, r *http.Request) {
}

func (h *GatewayApiHandler) deleteCategory(w http.ResponseWriter, r *http.Request) {
}
