package handler

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
)

func (h *GatewayApiHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	var payload productPb.Product

	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateProductPayload(&payload); len(errs) > 0 {
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.productServiceClient.CreateProduct(r.Context(), &productPb.CreateProductRequest{
		Product: &payload,
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusCreated, res)
}
