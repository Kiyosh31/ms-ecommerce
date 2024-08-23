package brand

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Update(ctx context.Context, brandToUpdate domain.BrandSchema) (*productPb.Brand, error) {
	return &productPb.Brand{}, nil
}
