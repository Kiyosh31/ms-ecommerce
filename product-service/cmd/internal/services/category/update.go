package category

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Update(ctx context.Context, categoryToUpdate domain.CategorySchema) (*productPb.Category, error) {
	return &productPb.Category{}, nil
}
