package user

import (
	"context"
	"net/http"
	"time"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
)

func (h *Handler) Reactivate(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("reactivate user request incoming: %v", customlogger.ReadRequestPayload(r))
	var payload userPb.ReactivateUserRequest

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("error reading payload: %v", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := ValidateReactivateUser(&payload); len(errs) > 0 {
		h.logger.Errorf("error validating payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.userServiceClient.ReactivateUser(ctx, &payload)
	if err != nil {
		h.logger.Errorf("error reactivating user: %v", err.Error())
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("reactivate user request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
