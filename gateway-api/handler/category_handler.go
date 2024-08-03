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
	categoryId := r.PathValue("categoryId")
	h.logger.Infof("get category request incoming: %v", categoryId)

	res, err := h.productServiceClient.GetCategory(r.Context(), &productPb.CategoryRequest{
		CategoryId: &categoryId,
	})
	if err != nil {
		h.logger.Errorf("failed to get category: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get category finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) updateCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("categoryId")
	h.logger.Infof("update category request incoming with id: %v and payload: %v", categoryId, customlogger.ReadRequestPayload(r))

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

	res, err := h.productServiceClient.UpdateCategory(r.Context(), &productPb.CategoryRequest{
		CategoryId: &categoryId,
		Category:   &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to update category: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("update category finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("categoryId")
	h.logger.Infof("delete category request incoming with id: %v", categoryId)

	res, err := h.productServiceClient.DeleteCategory(r.Context(), &productPb.CategoryRequest{
		CategoryId: &categoryId,
	})
	if err != nil {
		h.logger.Errorf("failed to delete category: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("delete category finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
