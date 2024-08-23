package product

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Update(ctx context.Context, productToUpdate domain.ProductSchema) (*productPb.Product, error) {
	return &productPb.Product{}, nil
}
