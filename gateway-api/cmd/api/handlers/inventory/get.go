package inventory

import (
	"context"
	"net/http"
	"time"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/inventory-service"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	inventoryId := r.PathValue("inventoryId")
	h.logger.Info("get inventory request incoming: %v", inventoryId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.inventoryServiceClient.SearchInventory(ctx, &inventoryPb.InventoryRequest{
		InventoryId: &inventoryId,
	})
	if err != nil {
		h.logger.Errorf("error trying to get inventory: %v", err.Error())
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get inventory request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
