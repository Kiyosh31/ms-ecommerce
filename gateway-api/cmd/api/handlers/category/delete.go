package category

import (
	"context"
	"net/http"
	"time"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
)

func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("categoryId")
	h.logger.Infof("delete category request incoming with id: %v", categoryId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.categoryServiceClient.DeleteCategory(ctx, &productPb.CategoryRequest{
		CategoryId: &categoryId,
	})
	if err != nil {
		h.logger.Errorf("failed to delete category: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("delete category finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
