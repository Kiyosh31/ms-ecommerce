package user

import (
	"context"
	"net/http"
	"time"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	h.logger.Infof("update user request incoming with id: %v and body: %v", userId, customlogger.ReadRequestPayload(r))

	var payload userPb.User

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("error reading request payload: %v", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := ValidateUserPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("error validating request payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.userServiceClient.UpdateUser(ctx, &userPb.UserRequest{
		UserId: &userId,
		User:   &payload,
	})
	if err != nil {
		h.logger.Errorf("error updating user: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("update user request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
