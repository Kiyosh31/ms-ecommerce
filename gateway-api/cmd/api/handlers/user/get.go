package user

import (
	"context"
	"net/http"
	"time"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	h.logger.Info("get user request incoming: %v", userId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.userServiceClient.GetUser(ctx, &userPb.UserRequest{
		UserId: &userId,
	})
	if err != nil {
		h.logger.Errorf("error trying to get user: %v", err.Error())
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get user request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
