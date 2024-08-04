package handler

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/inventory-service"
)

func (h *GatewayApiHandler) createInventory(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create product request incoming: %v", customlogger.ReadRequestPayload(r))
	var payload inventoryPb.Inventory

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read inventory payload: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateInventoryPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate inventory payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.inventoryServiceClient.CreateInventory(r.Context(), &inventoryPb.InventoryRequest{
		Inventory: &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to create inventory: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create inventory finished: %v", res)
	utils.WriteResponse(w, http.StatusCreated, res)
}

func (h *GatewayApiHandler) getInventory(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("get inventory request incoming: %v", customlogger.ReadRequestPayload(r))
	inventoryId := r.PathValue("inventoryId")

	res, err := h.inventoryServiceClient.GetInventory(r.Context(), &inventoryPb.InventoryRequest{
		InventoryId: &inventoryId,
	})
	if err != nil {
		h.logger.Errorf("failed to get inventory: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get inventory request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) updateInventory(w http.ResponseWriter, r *http.Request) {
	inventoryId := r.PathValue("inventoryId")
	h.logger.Infof("update inventory request incoming with id: %v and body: %v", inventoryId, customlogger.ReadRequestPayload(r))

	var payload inventoryPb.Inventory

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read inventory payload: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateInventoryPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate inventory payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.inventoryServiceClient.UpdateInventory(r.Context(), &inventoryPb.InventoryRequest{
		InventoryId: &inventoryId,
		Inventory:   &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to update inventory: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("update inventory request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
