package category

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) CreateCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	h.logger.Infof("create category request incoming: %v", in)

	// translate request
	// validate request
	newCategory := domain.CategorySchema{
		Name:        in.GetCategory().GetName(),
		Description: in.GetCategory().GetDescription(),
	}

	//consume service
	res, err := h.categoryService.Create(ctx, newCategory)
	if err != nil {
		return &productPb.CategoryResponse{}, err
	}

	// translate response
	return &productPb.CategoryResponse{
		Message:  "category created successfully",
		Category: []*productPb.Category{res},
	}, nil
}
