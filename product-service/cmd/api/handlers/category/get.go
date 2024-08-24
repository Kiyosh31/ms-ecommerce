package category

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) GetCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	// translate request
	// validate request

	//consume service
	res, err := h.categoryService.Get(ctx, in.GetCategoryId())
	if err != nil {
		return &productPb.CategoryResponse{}, err
	}

	// translate response
	return &productPb.CategoryResponse{
		Message:  "Category found",
		Category: []*productPb.Category{res},
	}, nil
}
