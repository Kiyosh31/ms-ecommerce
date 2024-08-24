package brand

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) DeleteBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	h.logger.Infof("create brand request incoming: %v", in)

	// translate request
	// validate request

	//consume service
	res, err := h.brandService.Delete(ctx, in.GetBrandId())
	if err != nil {
		return &productPb.BrandResponse{}, err
	}

	message := ""
	if res {
		message = "brand deleted successfully"
	} else {
		message = "couldnt delete"
	}

	// translate response
	return &productPb.BrandResponse{
		Message: message,
	}, nil
}
