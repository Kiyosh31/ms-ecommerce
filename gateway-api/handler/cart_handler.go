package handler

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	cartPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/cart-service"
)

func (h *GatewayApiHandler) createCart(w http.ResponseWriter, r *http.Request) {
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

func (h *GatewayApiHandler) getCart(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	cartId := r.PathValue("cartId")

	res, err := h.cartProductClient.GetCart(r.Context(), &cartPb.GetCartRequest{
		UserId: userId,
		CartId: cartId,
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) getAllCarts(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	res, err := h.cartProductClient.GetAllCarts(r.Context(), &cartPb.GetCartsRequest{
		UserId: userId,
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}
