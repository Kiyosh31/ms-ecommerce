package category

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) DeleteProduct(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	return &productPb.CategoryResponse{}, nil
}
