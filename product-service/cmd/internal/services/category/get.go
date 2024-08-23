package category

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Get(ctx context.Context, id string) (*productPb.Category, error) {
	return &productPb.Category{}, nil
}
