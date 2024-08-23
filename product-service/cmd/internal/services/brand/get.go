package brand

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Get(ctx context.Context, id string) (*productPb.Brand, error) {
	return &productPb.Brand{}, nil
}
