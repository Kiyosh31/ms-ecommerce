package brand

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) GetBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	// translate request
	// validate request

	//consume service
	res, err := h.brandService.Get(ctx, in.GetBrandId())
	if err != nil {
		return &productPb.BrandResponse{}, err
	}

	// translate response
	return &productPb.BrandResponse{
		Message: "Brand found",
		Brand:   []*productPb.Brand{res},
	}, nil
}
