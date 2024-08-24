package brand

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) CreateBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	h.logger.Infof("create brand request incoming: %v", in)

	// translate request
	// validate request
	newBrand := domain.BrandSchema{
		Name:        in.GetBrand().GetName(),
		Description: in.GetBrand().GetDescription(),
	}

	//consume service
	res, err := h.brandService.Create(ctx, newBrand)
	if err != nil {
		return &productPb.BrandResponse{}, err
	}

	// translate response
	return &productPb.BrandResponse{
		Message: "brand created successfully",
		Brand:   []*productPb.Brand{res},
	}, nil
}
