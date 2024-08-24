package brand

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (h *Handler) UpdateBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	h.logger.Infof("update brand request incoming: %v", in)

	// translate request
	brandId, err := database.GetMongoId(in.GetBrandId())
	if err != nil {
		return &productPb.BrandResponse{}, err
	}
	// validate request
	updateBrand := domain.BrandSchema{
		ID:          brandId,
		Name:        in.GetBrand().GetName(),
		Description: in.GetBrand().GetDescription(),
	}

	//consume service
	res, err := h.brandService.Update(ctx, updateBrand)
	if err != nil {
		return &productPb.BrandResponse{}, err
	}

	// translate response
	return &productPb.BrandResponse{
		Message: "brand updated successfully",
		Brand:   []*productPb.Brand{res},
	}, nil
}
