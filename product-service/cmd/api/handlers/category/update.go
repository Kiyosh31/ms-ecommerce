package category

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) UpdateCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	h.logger.Infof("update category request incoming: %v", in)

	// translate request
	categoryId, err := database.GetMongoId(in.GetCategoryId())
	if err != nil {
		return &productPb.CategoryResponse{}, err
	}
	// validate request
	updateCategory := domain.CategorySchema{
		ID:          categoryId,
		Name:        in.GetCategory().GetName(),
		Description: in.GetCategory().GetDescription(),
	}

	//consume service
	res, err := h.categoryService.Update(ctx, updateCategory)
	if err != nil {
		return &productPb.CategoryResponse{}, err
	}

	// translate response
	return &productPb.CategoryResponse{
		Message:  "category created successfully",
		Category: []*productPb.Category{res},
	}, nil
}
