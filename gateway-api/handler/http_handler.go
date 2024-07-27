package handler

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
)

type GatewayApiHandler struct {
	userServiceClient userPb.UserServiceClient
}

func NewHandler(userServiceClient userPb.UserServiceClient) *GatewayApiHandler {
	return &GatewayApiHandler{
		userServiceClient: userServiceClient,
	}
}

func (h *GatewayApiHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/user", h.createUser)
}

func (h *GatewayApiHandler) createUser(w http.ResponseWriter, r *http.Request) {

	var payload userPb.User
	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateUserPayload(&payload); len(errs) > 0 {
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.userServiceClient.CreateUser(r.Context(), &userPb.CreateUserRequest{
		User: &payload,
	})
	if err != nil {
		utils.ManageRpcErrors(err, w)
		return
	}

	utils.WriteJSON(w, http.StatusOK, res)
}
