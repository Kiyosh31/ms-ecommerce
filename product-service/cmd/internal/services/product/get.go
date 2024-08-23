package product

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Get(ctx context.Context, id string) (*productPb.Product, error) {
	return &productPb.Product{}, nil
}
