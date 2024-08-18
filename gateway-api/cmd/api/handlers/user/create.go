package user

import (
	"context"
	"net/http"
	"time"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create user request incoming: %v", customlogger.ReadRequestPayload(r))

	var payload userPb.User

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("error reading req json: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := ValidateUserPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("error validating req json")
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &userPb.UserRequest{
		User: &payload,
	}

	res, err := h.userServiceClient.CreateUser(ctx, req)
	if err != nil {
		h.logger.Errorf("error creating user: %v", err.Error())
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create user request finshed: %v", res)
	utils.WriteResponse(w, http.StatusCreated, res)
}
