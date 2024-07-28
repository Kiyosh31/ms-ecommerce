package handler

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
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
	// User endpoints
	mux.HandleFunc("POST /api/v1/user", h.createUser)
	mux.HandleFunc("GET /api/v1/user/{userId}", h.getUser)
	mux.HandleFunc("PUT /api/v1/user/{userId}", h.updateUser)
	mux.HandleFunc("DELETE /api/v1/user/{userId}", h.deleteUser)
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

	res, err := h.userServiceClient.CreateUser(r.Context(), mapCreateUserRequestToPb(&payload))
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusCreated, res)
}

func (h *GatewayApiHandler) getUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	res, err := h.userServiceClient.GetUser(r.Context(), &userPb.GetUserRequest{
		UserId: userId,
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) updateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *GatewayApiHandler) deleteUser(w http.ResponseWriter, r *http.Request) {

}
