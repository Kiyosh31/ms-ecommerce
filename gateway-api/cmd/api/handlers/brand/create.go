package brand

import (
	"context"
	"net/http"
	"time"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	"github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/api/handlers/product"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
)

func (h *Handler) CreateBrand(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create brand request incoming: %v", customlogger.ReadRequestPayload(r))

	var payload productPb.Brand

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read brand payload: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := product.ValidateBrandPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate brand payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.brandServiceClient.CreateBrand(ctx, &productPb.BrandRequest{
		Brand: &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to create category: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create brand finished: %v", res)
	utils.WriteResponse(w, http.StatusCreated, res)
}
