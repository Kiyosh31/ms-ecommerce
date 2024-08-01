package handler

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	cartPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/cart-service"
)

func (h *GatewayApiHandler) createCart(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create cart request incoming: %v", customlogger.ReadRequestPayload(r))
	var payload cartPb.Cart

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read cart payload: %v", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateCartPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate cart payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.cartProductClient.CreateCart(r.Context(), &cartPb.CreateCartRequest{
		Cart: &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to create cart: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create cart request finished: %v", err)
	utils.WriteResponse(w, http.StatusCreated, res)
}

func (h *GatewayApiHandler) getCart(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	cartId := r.PathValue("cartId")
	h.logger.Infof("get cart incoming request with userId: %v and cartId: %v", userId, cartId)

	res, err := h.cartProductClient.GetCart(r.Context(), &cartPb.GetCartRequest{
		UserId: userId,
		CartId: cartId,
	})
	if err != nil {
		h.logger.Errorf("failed to get cart: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get cart request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) getAllCarts(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	h.logger.Infof("get all carts request incoming: %v", userId)

	res, err := h.cartProductClient.GetAllCarts(r.Context(), &cartPb.GetCartsRequest{
		UserId: userId,
	})
	if err != nil {
		h.logger.Errorf("failed to get all carts: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get all carts request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
