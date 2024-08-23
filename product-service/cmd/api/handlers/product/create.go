package product

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) CreateProduct(ctx context.Context, in *productPb.ProductRequest) (*productPb.ProductResponse, error) {

	return &productPb.ProductResponse{}, nil
}
