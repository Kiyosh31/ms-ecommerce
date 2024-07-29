package handler

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
)

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
	userId := r.PathValue("userId")

	var payload userPb.User

	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateUserPayload(&payload); len(errs) > 0 {
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.userServiceClient.UpdateUser(r.Context(), &userPb.UpdateUserRequest{
		User: &userPb.User{
			Id:        userId,
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			BirthDate: payload.BirthDate,
			Email:     payload.Email,
			Password:  payload.Password,
			Cards:     payload.Cards,
			Addresses: payload.Addresses,
			IsActive:  true,
		},
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	if userId == "" {
		utils.WriteError(w, http.StatusBadRequest, "userId missing")
		return
	}

	res, err := h.userServiceClient.DeleteUser(r.Context(), &userPb.DeleteUserRequest{
		UserId: userId,
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) reactivateUser(w http.ResponseWriter, r *http.Request) {
	var payload userPb.ReactivarUserRequest

	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateReactivateUser(&payload); len(errs) > 0 {
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.userServiceClient.ReactivateUser(r.Context(), &payload)
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}
