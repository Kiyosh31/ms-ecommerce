package user

import (
	"context"
	"net/http"
	"time"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
)

func (h *Handler) Deactivate(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	h.logger.Info("delete user request incoming: %v", userId)
	if userId == "" {
		h.logger.Error("error no userId")
		utils.WriteError(w, http.StatusBadRequest, "userId missing")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.userServiceClient.DeactivateUser(ctx, &userPb.UserRequest{
		UserId: &userId,
	})
	if err != nil {
		h.logger.Errorf("error deleting user: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("delete user request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
