package handler

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
)

func (h *GatewayApiHandler) createProduct(w http.ResponseWriter, r *http.Request) {

	res, err := h.productServiceClient.CreateProduct(r.Context(), &productPb.CreateProductRequest{
		Product: &productPb.Product{
			Id:          "1",
			Price:       123.44,
			Description: "asdasda",
		},
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusCreated, res)
}
