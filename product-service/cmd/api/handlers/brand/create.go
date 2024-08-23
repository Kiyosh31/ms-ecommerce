package brand

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) CreateProduct(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {

	return &productPb.BrandResponse{}, nil
}
