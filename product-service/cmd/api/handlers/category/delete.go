package category

import (
	"context"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) DeleteCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	h.logger.Infof("create category request incoming: %v", in)

	// translate request
	// validate request

	//consume service
	res, err := h.categoryService.Delete(ctx, in.GetCategoryId())
	if err != nil {
		return &productPb.CategoryResponse{}, err
	}

	message := ""
	if res {
		message = "category deleted successfully"
	} else {
		message = "couldnt delete"
	}

	// translate response
	return &productPb.CategoryResponse{
		Message: message,
	}, nil

}
