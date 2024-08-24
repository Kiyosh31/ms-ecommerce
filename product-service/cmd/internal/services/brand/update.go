package brand

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Update(ctx context.Context, brandToUpdate domain.BrandSchema) (*productPb.Brand, error) {
	s.logger.Infof("update brand: %v", brandToUpdate)

	// check if category exists
	_, exist, err := s.repository.Get(ctx, brandToUpdate.ID)
	if err != nil {
		return &productPb.Brand{}, err
	}
	if !exist {
		return &productPb.Brand{}, errors.New("brand not found")
	}

	_, err = s.repository.Update(ctx, brandToUpdate)
	if err != nil {
		s.logger.Error(err)
		return &productPb.Brand{}, err
	}

	s.logger.Infof("update brand finished with: %v", brandToUpdate)
	return &productPb.Brand{
		Id:          brandToUpdate.ID.Hex(),
		Name:        brandToUpdate.Name,
		Description: brandToUpdate.Description,
	}, nil
}
