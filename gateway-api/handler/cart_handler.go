package handler

import (
	"log"
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	cartPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/cart-service"
)

func (h *GatewayApiHandler) createCart(w http.ResponseWriter, r *http.Request) {
	log.Println("llegue aqui")
	var payload cartPb.Cart

	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateCartPayload(&payload); len(errs) > 0 {
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.cartProductClient.CreateCart(r.Context(), &cartPb.CreateCartRequest{
		Cart: &payload,
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusCreated, res)
}
