package handler

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
)

func (h *GatewayApiHandler) createUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create user request incoming: %v", customlogger.ReadRequestPayload(r))
	var payload userPb.User

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("error reading req json: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateUserPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("error validating req json")
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.userServiceClient.CreateUser(r.Context(), &userPb.CreateUserRequest{
		User: &payload,
	})
	if err != nil {
		h.logger.Errorf("error validating req json: %v", err.Error())
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create user request finshed: %v", res)
	utils.WriteResponse(w, http.StatusCreated, res)
}

func (h *GatewayApiHandler) getUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	h.logger.Info("get user request incoming: %v", userId)

	res, err := h.userServiceClient.GetUser(r.Context(), &userPb.GetUserRequest{
		UserId: userId,
	})
	if err != nil {
		h.logger.Errorf("error trying to get user: %v", err.Error())
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get user request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	h.logger.Infof("update user request incoming with id: %v and body: %v", userId, customlogger.ReadRequestPayload(r))

	var payload userPb.User

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("error reading request payload: %v", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateUserPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("error validating request payload: %v", errs)
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
		h.logger.Errorf("error updating user: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("update user request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	h.logger.Info("delete user request incoming: %v", userId)
	if userId == "" {
		h.logger.Error("error no userId")
		utils.WriteError(w, http.StatusBadRequest, "userId missing")
		return
	}

	res, err := h.userServiceClient.DeleteUser(r.Context(), &userPb.DeleteUserRequest{
		UserId: userId,
	})
	if err != nil {
		h.logger.Errorf("error deleting user: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("delete user request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) reactivateUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("reactivate user request incoming: %v", customlogger.ReadRequestPayload(r))
	var payload userPb.ReactivarUserRequest

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("error reading payload: %v", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateReactivateUser(&payload); len(errs) > 0 {
		h.logger.Errorf("error validating payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.userServiceClient.ReactivateUser(r.Context(), &payload)
	if err != nil {
		h.logger.Errorf("error reactivating user: %v", err.Error())
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("reactivate user request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
